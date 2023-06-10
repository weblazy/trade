package spot_trade_engine

import (
	"sync"
	"time"

	"trade/https/trade_api/def"
	"trade/pkg/cache"
	"trade/pkg/sortedset"
)

type TradeEngine struct {
	WorkerMap map[string]*Worker // key:symbol
	WaitGroup sync.WaitGroup
}

func NewTradeEngine() (*TradeEngine, error) {
	trade := &TradeEngine{
		WorkerMap: make(map[string]*Worker),
	}
	return trade, nil
}

func (t *TradeEngine) AddSymbols(symbols []string) {
	for _, symbol := range symbols {
		if _, ok := t.WorkerMap[symbol]; !ok {
			worker := &Worker{
				OrderChan: make(chan *def.Order, 100),
				BuyQueue:  sortedset.NewSortedSet(),
				SellQueue: sortedset.NewSortedSet(),
			}
			t.WorkerMap[symbol] = worker
			cache.SaveSymbol(symbol)
			t.WaitGroup.Add(1)
			go func(worker *Worker) {
				defer t.WaitGroup.Done()
				worker.Run()
			}(worker)
			orderList := cache.GetOrderList(symbol)
			for k := range orderList {
				worker.OrderChan <- orderList[k]
			}
		}
	}
}

func (t *TradeEngine) DeleteSymbols(symbols []string) {
	for _, symbol := range symbols {
		if worker, ok := t.WorkerMap[symbol]; ok {
			worker.IsClosed.Store(true)
			delete(t.WorkerMap, symbol)
			time.Sleep(10 * time.Second)
			close(worker.OrderChan)
		}
	}
}

func (t *TradeEngine) OnpenTrade(symbols []string) {
	for _, symbol := range symbols {
		if worker, ok := t.WorkerMap[symbol]; ok {
			worker.IsClosed.Store(false)
		}
	}
}

func (t *TradeEngine) onpenTrade(symbol string) error {
	defer t.WaitGroup.Done()
	worker, ok := t.WorkerMap[symbol]
	if !ok {
		return WorkerNotFoundErr
	}
	worker.IsClosed.Store(false)
	for {
		select {
		case order, ok := <-worker.OrderChan:
			if !ok {
				delete(t.WorkerMap, symbol)
				cache.Clear(symbol)
				return OrderChanCloseErr
			}
			worker.CreateOrder(order)
		case orderNo, ok := <-worker.OrderNoChan:
			if !ok {
				delete(t.WorkerMap, symbol)
				cache.Clear(symbol)
				return OrderChanCloseErr
			}
			worker.CancelOrder(orderNo)
		}
	}
}

func (t *TradeEngine) CloseTrade(symbol string) error {
	worker, ok := t.WorkerMap[symbol]
	if !ok {
		return WorkerNotFoundErr
	}
	worker.IsClosed.Store(true)
	return nil
}

func (t *TradeEngine) CreateOrder(order *def.Order) error {
	worker := t.WorkerMap[order.Symbol]
	if worker == nil {
		return WorkerNotFoundErr
	}
	if worker.IsClosed.Load() {
		return TradeClosedErr
	}
	worker.OrderChan <- order
	return nil
}

func (t *TradeEngine) CancelOrder(symbol, orderNo string) error {
	if t.WorkerMap[symbol] == nil {
		return WorkerNotFoundErr
	}
	t.WorkerMap[symbol].OrderNoChan <- orderNo
	return nil
}

func (t *TradeEngine) Close() {
	t.WaitGroup.Wait()
}
