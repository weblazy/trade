package def

import "github.com/shopspring/decimal"

type Action string

func (a Action) String() string {
	return string(a)
}

const (
	ActionCreate Action = "ActionCreate"
	ActionCancel Action = "ActionCancel"
)

type Side string

func (a Side) String() string {
	return string(a)
}

const (
	Buy  Side = "Buy"
	Sell Side = "Sell"
)

type Order struct {
	UserId      int64           `json:"user_id"` // 用户id
	Pair        string          `json:"Pair"`    // 交易对
	Price       decimal.Decimal `json:"price"`   // 价格
	OrderId     string          `json:"order_id"`
	OrderNo     string          `json:"order_no"`
	Symbol      string          `json:"symbol"`
	Action      Action          `json:"action"`
	Side        Side            `json:"side"`
	OrderType   OrderType       `json:"order_type"`
	LastAmount  decimal.Decimal `json:"last_amount"`  // 剩下的数量
	TotalAmount decimal.Decimal `json:"total_amount"` // 总数量
	TimeInForce string          `json:"f"`            // 订单有效时间,type为limit时才生效 GTC/IOC/FOK

}
type Trade struct {
	Id               string          `json:"i"`  // 成交单id
	Pair             string          `json:"P"`  // 交易对
	MakerId          string          `json:"mi"` // maker订单id
	TakerId          string          `json:"ti"` // taker订单id
	MakerUser        int64           `json:"mu"` // maker用户id
	TakerUser        int64           `json:"tu"` // taker用户id
	Price            decimal.Decimal `json:"p"`  // 成交价
	Amount           decimal.Decimal `json:"a"`  // 成交数量
	TakerOrderSide   Side            `json:"s"`  // taker订单方向 buy/sell
	TakerOrderType   OrderType       `json:"t"`  // taker订单类型 limit/market/cancel
	TakerTimeInForce string          `json:"f"`  // taker订单有效时间,type为limit时才生效 GTC/IOC/FOK
	Ts               int64           `json:"ts"` // 成交时间
}

type OrderType string

const (
	TypeLimit  OrderType = "TypeLimit"
	TypeMarket OrderType = "TypeMarket"
)

func (a OrderType) String() string {
	return string(a)
}

type FutrueOrder struct {
	UserId      int64           `json:"user_id"` // 用户id
	Pair        string          `json:"Pair"`    // 交易对
	Price       decimal.Decimal `json:"price"`   // 价格
	OrderId     string          `json:"order_id"`
	OrderNo     string          `json:"order_no"`
	Symbol      string          `json:"symbol"`
	Action      Action          `json:"action"`
	Side        Side            `json:"side"`
	OrderType   OrderType       `json:"order_type"`
	LastAmount  decimal.Decimal `json:"last_amount"`  // 剩下的数量
	TotalAmount decimal.Decimal `json:"total_amount"` // 总数量
	TimeInForce string          `json:"f"`            // 订单有效时间,type为limit时才生效 GTC/IOC/FOK

	Margin           decimal.Decimal `json:"margin"`   //保证金
	Value            decimal.Decimal `json:"value"`    //价值
	Leverage         decimal.Decimal `json:"leverage"` //倍数
	LiquidationPrice decimal.Decimal `json:"value"`    //强平价
	MarkPrice        decimal.Decimal `json:"value"`    //标价
	EntryPrice       decimal.Decimal `json:"value"`    //入场价

}

type FutrueTrade struct {
	UserId      int64           `json:"user_id"` // 用户id
	Pair        string          `json:"Pair"`    // 交易对
	Price       decimal.Decimal `json:"price"`   // 价格
	OrderId     string          `json:"order_id"`
	OrderNo     string          `json:"order_no"`
	Symbol      string          `json:"symbol"`
	Action      Action          `json:"action"`
	Side        Side            `json:"side"`
	OrderType   OrderType       `json:"order_type"`
	LastAmount  decimal.Decimal `json:"last_amount"`  // 剩下的数量
	TotalAmount decimal.Decimal `json:"total_amount"` // 总数量
	TimeInForce string          `json:"f"`            // 订单有效时间,type为limit时才生效 GTC/IOC/FOK

	Margin           decimal.Decimal `json:"margin"`   //保证金
	Value            decimal.Decimal `json:"value"`    //价值
	Leverage         decimal.Decimal `json:"leverage"` //倍数
	LiquidationPrice decimal.Decimal `json:"value"`    //强平价
	MarkPrice        decimal.Decimal `json:"value"`    //标价
	EntryPrice       decimal.Decimal `json:"value"`    //入场价

}
