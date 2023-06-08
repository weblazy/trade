package trade_engine

import (
	"strconv"
	"time"
	"trade/https/trade_api/def"
	"trade/pkg/cache"
	"trade/pkg/sortedset"

	"github.com/shopspring/decimal"
)

type Worker struct {
	OrderChan chan *def.Order
	OrderMap  map[string]*def.Order
	BuyQueue  *sortedset.SortedSet //Element.Value->*Data
	SellQueue *sortedset.SortedSet //Element.Value->*Data
}

type Data struct {
	Key   string
	Value *def.Order
}

func (w *Worker) CreateOrder(order *def.Order) {
	if order.Side == def.Buy {
		switch order.OrderType {
		case def.TypeLimit:
			w.AddLimitBuy(order)
		case def.TypeMarket:
			w.AddMarketSell(order)
		}
	} else {
		switch order.OrderType {
		case def.TypeLimit:
			w.AddLimitSell(order)
		case def.TypeMarket:
			w.AddMarketSell(order)
		}
	}

}

func (w *Worker) AddLimitBuy(order *def.Order) {
	trades := make([]*def.Trade, 0)
	for {
		// 获取卖方队列最小价格订单
		e := w.SellQueue.GetFirst()
		if e == nil {
			break
		}
		sellOrder := w.OrderMap[e.Member]
		if sellOrder == nil {
			continue
		}
		// 订单金额小于卖单价格时无法成交退出循环
		if order.Price.LessThan(sellOrder.Price) {
			break
		}
		// 剩余数量为0时退出循环
		if order.LastAmount.LessThanOrEqual(decimal.Zero) {
			break
		}
		trade := &def.Trade{
			Id:               GenTradeId(),
			Pair:             order.Pair,
			MakerId:          sellOrder.OrderNo,
			TakerId:          order.OrderNo,
			MakerUser:        sellOrder.UserId,
			TakerUser:        order.UserId,
			Price:            sellOrder.Price,
			TakerOrderSide:   order.Side,
			TakerOrderType:   order.OrderType,
			TakerTimeInForce: order.TimeInForce,
			Ts:               NowUnixMilli(),
		}
		// 卖单数量大于等于买单,taker全部成交
		if sellOrder.LastAmount.GreaterThanOrEqual(order.LastAmount) {
			trade.Amount = order.LastAmount
		} else {
			// 卖单数量小于买单,maker全部分成交
			trade.Amount = sellOrder.LastAmount
		}
		trades = append(trades, trade)
		sellOrder.LastAmount = sellOrder.LastAmount.Sub(trade.Amount)
		order.LastAmount = order.LastAmount.Sub(trade.Amount)
		// 判断sellOrder剩余数量 <= 0
		if sellOrder.LastAmount.LessThanOrEqual(decimal.Zero) {
			w.SellQueue.Remove(sellOrder.OrderNo)
			delete(w.OrderMap, sellOrder.OrderNo)
		}
	}

	// 判断order是否完全成交
	if order.LastAmount.GreaterThan(decimal.Zero) {
		score, _ := order.Price.Float64()
		w.BuyQueue.Add(order.OrderNo, score)
		w.OrderMap[order.OrderNo] = order
	}

	if len(trades) > 0 {
		w.PushTrades(trades)
	}
}

func (w *Worker) AddLimitSell(order *def.Order) {
	trades := make([]*def.Trade, 0)
	for {
		// 获取买方队列最大价格订单
		e := w.BuyQueue.GetLast()
		if e == nil {
			break
		}
		buyOrder := w.OrderMap[e.Member]
		if buyOrder == nil {
			continue
		}
		// 订单金额大于买单价格时无法成交退出循环
		if order.Price.GreaterThan(buyOrder.Price) {
			break
		}
		// 剩余数量为0时退出循环
		if order.LastAmount.LessThanOrEqual(decimal.Zero) {
			break
		}
		trade := &def.Trade{
			Id:               GenTradeId(),
			Pair:             order.Pair,
			MakerId:          buyOrder.OrderNo,
			TakerId:          order.OrderNo,
			MakerUser:        buyOrder.UserId,
			TakerUser:        order.UserId,
			Price:            buyOrder.Price,
			TakerOrderSide:   order.Side,
			TakerOrderType:   order.OrderType,
			TakerTimeInForce: order.TimeInForce,
			Ts:               NowUnixMilli(),
		}
		// maker数量大于等于taker,taker全部成交
		if buyOrder.LastAmount.GreaterThanOrEqual(order.LastAmount) {
			trade.Amount = order.LastAmount
		} else {
			// maker数量小于taker,maker全部分成交
			trade.Amount = buyOrder.LastAmount
		}
		trades = append(trades, trade)
		buyOrder.LastAmount = buyOrder.LastAmount.Sub(trade.Amount)
		order.LastAmount = order.LastAmount.Sub(trade.Amount)
		// 判断maker剩余数量 <= 0
		if buyOrder.LastAmount.LessThanOrEqual(decimal.Zero) {
			w.BuyQueue.Remove(buyOrder.OrderNo)
			delete(w.OrderMap, buyOrder.OrderNo)
		}
	}

	// 判断order是否完全成交
	if order.LastAmount.GreaterThan(decimal.Zero) {
		score, _ := order.Price.Float64()
		w.SellQueue.Add(order.OrderNo, score)
		w.OrderMap[order.OrderNo] = order
	}

	if len(trades) > 0 {
		w.PushTrades(trades)
	}
}
func (w *Worker) AddMarketBuy(order *def.Order) {
	trades := make([]*def.Trade, 0)
	for {
		// 获取卖方队列最小价格订单
		e := w.SellQueue.GetFirst()
		if e == nil {
			break
		}
		sellOrder := w.OrderMap[e.Member]
		if sellOrder == nil {
			continue
		}
		// 剩余数量为0时退出循环
		if order.LastAmount.LessThanOrEqual(decimal.Zero) {
			break
		}
		trade := &def.Trade{
			Id:               GenTradeId(),
			Pair:             order.Pair,
			MakerId:          sellOrder.OrderNo,
			TakerId:          order.OrderNo,
			MakerUser:        sellOrder.UserId,
			TakerUser:        order.UserId,
			Price:            sellOrder.Price,
			TakerOrderSide:   order.Side,
			TakerOrderType:   order.OrderType,
			TakerTimeInForce: order.TimeInForce,
			Ts:               NowUnixMilli(),
		}
		// 卖单数量大于等于买单,taker全部成交
		if sellOrder.LastAmount.GreaterThanOrEqual(order.LastAmount) {
			trade.Amount = order.LastAmount
		} else {
			// 卖单数量小于买单,maker全部分成交
			trade.Amount = sellOrder.LastAmount
		}
		trades = append(trades, trade)
		sellOrder.LastAmount = sellOrder.LastAmount.Sub(trade.Amount)
		order.LastAmount = order.LastAmount.Sub(trade.Amount)
		// 判断sellOrder剩余数量 <= 0
		if sellOrder.LastAmount.LessThanOrEqual(decimal.Zero) {
			w.SellQueue.Remove(sellOrder.OrderNo)
			delete(w.OrderMap, sellOrder.OrderNo)
		}
	}

	if len(trades) > 0 {
		w.PushTrades(trades)
	}
}
func (w *Worker) AddMarketSell(order *def.Order) {
	trades := make([]*def.Trade, 0)
	for {
		// 获取买方队列最大价格订单
		e := w.BuyQueue.GetLast()
		if e == nil {
			break
		}
		buyOrder := w.OrderMap[e.Member]
		if buyOrder == nil {
			continue
		}

		// 剩余数量为0时退出循环
		if order.LastAmount.LessThanOrEqual(decimal.Zero) {
			break
		}
		trade := &def.Trade{
			Id:               GenTradeId(),
			Pair:             order.Pair,
			MakerId:          buyOrder.OrderNo,
			TakerId:          order.OrderNo,
			MakerUser:        buyOrder.UserId,
			TakerUser:        order.UserId,
			Price:            buyOrder.Price,
			TakerOrderSide:   order.Side,
			TakerOrderType:   order.OrderType,
			TakerTimeInForce: order.TimeInForce,
			Ts:               NowUnixMilli(),
		}
		// maker数量大于等于taker,taker全部成交
		if buyOrder.LastAmount.GreaterThanOrEqual(order.LastAmount) {
			trade.Amount = order.LastAmount
		} else {
			// maker数量小于taker,maker全部分成交
			trade.Amount = buyOrder.LastAmount
		}
		trades = append(trades, trade)
		buyOrder.LastAmount = buyOrder.LastAmount.Sub(trade.Amount)
		order.LastAmount = order.LastAmount.Sub(trade.Amount)
		// 判断maker剩余数量 <= 0
		if buyOrder.LastAmount.LessThanOrEqual(decimal.Zero) {
			w.BuyQueue.Remove(buyOrder.OrderNo)
			delete(w.OrderMap, buyOrder.OrderNo)
		}
	}

	if len(trades) > 0 {
		w.PushTrades(trades)
	}
}

func (w *Worker) CancelOrder(orderNo string) {
	order, ok := w.OrderMap[orderNo]
	if ok {
		switch order.Side {
		case def.Buy:
			w.BuyQueue.Remove(order.OrderNo)
		case def.Sell:
			w.SellQueue.Remove(order.OrderNo)
		}
		delete(w.OrderMap, orderNo)
		cache.RemoveOrder(order)
	}
}

// GenTradeId 生成成交单id
func GenTradeId() string {
	return strconv.Itoa(int(time.Now().UnixMilli()))
}
func NowUnixMilli() int64 {
	return time.Now().UnixMilli()
}

func (w *Worker) PushTrades(trades []*def.Trade) {
}
