package cache

import (
	"trade/https/trade_api/def"

	"github.com/shopspring/decimal"
)

func GetPrice(symbol string) decimal.Decimal {
	return decimal.Decimal{}
}

func GetOrderIdsWithAction(symbol string) {

}

func SaveSymbol(symbol string) {

}
func SavePrice(symbol string, price decimal.Decimal) {

}

func Clear(symbol string) {

}

func OrderExist(symbol string, orderId string, action def.Action) {

}
func RemoveOrder(order *def.Order) {

}
