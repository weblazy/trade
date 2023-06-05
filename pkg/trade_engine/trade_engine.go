package trade_engine

import (
	"sync"
	"time"

	"trade/https/trade_api/def"
	"trade/pkg/cache"
	"trade/pkg/sortedset"

	"github.com/pingcap/log"
	"github.com/shopspring/decimal"
	"gomod.sunmi.com/gomoddepend/sdk/order"
)

type TradeEngine struct {
	WorkerMap map[string]*Worker // key:symbol
	WaitGroup sync.WaitGroup
}

func NewTradeEngine() (*TradeEngine, error) {

	return &TradeEngine{
		WorkerMap: make(map[string]*Worker),
	}, nil
}

func (t *TradeEngine) Run(symbols []string) {
	for _, v := range symbols {
		if _, ok := t.WorkerMap[v]; !ok {
			t.WorkerMap[v] = &Worker{
				OrderChan: make(chan *def.Order, 100),
				BuyQueue:  sortedset.NewSortedSet(),
				SellQueue: sortedset.NewSortedSet(),
			}
		}
	}
}

func (t *TradeEngine) AddSymbols(symbol string, price decimal.Decimal) {
	if _, ok := t.WorkerMap[symbol]; !ok {
		t.WorkerMap[symbol] = &Worker{
			OrderChan: make(chan *order.Order, 100),
		}
	}
}

func (t *TradeEngine) OnpenTrade(symbol string, price decimal.Decimal) {
	cache.SaveSymbol(symbol)
	cache.SavePrice(symbol, price)
	t.WaitGroup.Add(1)
	go t.onpenTrade(symbol, price)
}

func (t *TradeEngine) onpenTrade(symbol string, price decimal.Decimal) error {
	defer t.WaitGroup.Done()
	worker, ok := t.WorkerMap[symbol]
	if !ok {
		return WorkerNotFoundErr
	}
	for {
		select {
		case order, ok := <-worker.OrderChan:
			if !ok {
				log.Info("engine %s is closed", symbol)
				delete(t.WorkerMap, symbol)
				cache.Clear(symbol)
				return OrderChanCloseErr
			}
			switch order.Action {
			case def.ActionCreate:
				worker.dealCreate(&order, book, &lastTradePrice)
			case def.ActionCancel:
				worker.dealCancel(order)
			}
		case task := <-tw.setChannel:
			t.setTask(&task)
		case key := <-tw.removeChannel:
			t.removeTask(key)
		case task := <-tw.moveChannel:
			t.moveTask(task)
		case fn := <-tw.drainChannel:
			t.drainAll(fn)
		case <-tw.stopChannel:
			t.ticker.Stop()
			return
		}
	}
}

func CloseTrade(symbol string, price decimal.Decimal) {

}

func Init(symbols []string) {
	for _, symbol := range symbols {
		price := cache.GetPrice(symbol)
		engine, _ := NewTradeEngine()

		orderIds := cache.GetOrderIdsWithAction(symbol)
		for _, orderId := range orderIds {
			mapOrder := cache.GetOrder(symbol, orderId)
			order := def.Order{}
			order.FromMap(mapOrder)
			engine.WorkerMap[order.Symbol].orderChan <- &order
		}
	}
}

func (t *TradeEngine) Dispatch(order *def.Order) error {
	if t.WorkerMap[order.Symbol] == nil {
		return WorkerNotFoundErr
	}

	if order.Action == def.ActionCreate {
		if cache.OrderExist(order.Symbol, order.OrderId, order.Action.String()) {
			return OrderExistErr
		}
	} else {
		if !cache.OrderExist(order.Symbol, order.OrderId, ActionCreate.String()) {
			return OrderNotFoundErr
		}
	}

	order.Timestamp = time.Now().UnixNano() / 1e3
	cache.SaveOrder(order.ToMap())
	t.WorkerMap[order.Symbol].orderChan <- order

	return nil
}

func (t *TradeEngine) Close() {
	t.WaitGroup.Wait()
}
