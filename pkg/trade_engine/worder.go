package trade_engine

import (
	"container/list"
	"trade/https/trade_api/def"
	"trade/pkg/cache"

	"github.com/shopspring/decimal"
)

type Worker struct {
	orderChan chan *def.Order
	BuyMap    map[string]*list.Element //key order_no
	BuyQueue  *list.List               //Element.Value->*Data
	SellMap   map[string]*list.Element //key order_no
	SellQueue *list.List               //Element.Value->*Data
}

type Data struct {
	Key   string
	Value *def.Order
}

func dealCreate(order *def.Order, lastTradePrice *decimal.Decimal) {
	switch order.OrderType {
	case def.TypeLimit:
		dealLimit(order, lastTradePrice)
	case def.TypeLimitIoc:
		dealLimitIoc(order, lastTradePrice)
	case def.TypeMarket:
		dealMarket(order, lastTradePrice)
	case def.TypeMarketTop5:
		dealMarketTop5(order, lastTradePrice)
	case def.TypeMarketTop10:
		dealMarketTop10(order, lastTradePrice)
	case def.TypeMarketOpponent:
		dealMarketOpponent(order, lastTradePrice)
	}
}
func dealLimit(order *def.Order, lastTradePrice *decimal.Decimal) {
}
func dealLimitIoc(order *def.Order, lastTradePrice *decimal.Decimal) {
}
func dealMarket(order *def.Order, lastTradePrice *decimal.Decimal) {
}
func dealMarketTop5(order *def.Order, lastTradePrice *decimal.Decimal) {
}
func dealMarketTop10(order *def.Order, lastTradePrice *decimal.Decimal) {
}
func dealMarketOpponent(order *def.Order, lastTradePrice *decimal.Decimal) {
}

func (w *Worker) dealCancel(order *def.Order) {
	switch order.Side {
	case def.Buy:
		e, ok := w.BuyMap[order.OrderNo]
		if ok {
			w.BuyQueue.Remove(e)
			delete(w.BuyMap, order.OrderNo)
		}
	case def.Sell:
		e, ok := w.SellMap[order.OrderNo]
		if ok {
			w.SellQueue.Remove(e)
			delete(w.SellMap, order.OrderNo)
		}
	}
	cache.RemoveOrder(order)
}
