package trade_engine

import (
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

func (w *Worker) dealCreate(order *def.Order, lastTradePrice *decimal.Decimal) {
	switch order.OrderType {
	case def.TypeLimit:
		w.dealLimit(order, lastTradePrice)
	case def.TypeLimitIoc:
		w.dealLimitIoc(order, lastTradePrice)
	case def.TypeMarket:
		w.dealMarket(order, lastTradePrice)
	case def.TypeMarketTop5:
		w.dealMarketTop5(order, lastTradePrice)
	case def.TypeMarketTop10:
		w.dealMarketTop10(order, lastTradePrice)
	case def.TypeMarketOpponent:
		w.dealMarketOpponent(order, lastTradePrice)
	}
}
func (w *Worker) dealLimit(order *def.Order, lastTradePrice *decimal.Decimal) {
	switch order.Side {
	case def.Buy:
		e := w.SellQueue.GetFirst()
		preOrder := w.OrderMap[e.Member]
		if order.Price.GreaterThanOrEqual(preOrder.Price) {

		}
	case def.Sell:

	}
}
func (w *Worker) dealLimitIoc(order *def.Order, lastTradePrice *decimal.Decimal) {
}
func (w *Worker) dealMarket(order *def.Order, lastTradePrice *decimal.Decimal) {
}
func (w *Worker) dealMarketTop5(order *def.Order, lastTradePrice *decimal.Decimal) {
}
func (w *Worker) dealMarketTop10(order *def.Order, lastTradePrice *decimal.Decimal) {
}
func (w *Worker) dealMarketOpponent(order *def.Order, lastTradePrice *decimal.Decimal) {
}

func (w *Worker) dealCancel(order *def.Order) {
	switch order.Side {
	case def.Buy:
		w.BuyQueue.Remove(order.OrderNo)
	case def.Sell:
		w.SellQueue.Remove(order.OrderNo)
	}
	cache.RemoveOrder(order)
}
