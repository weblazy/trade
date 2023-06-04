package def

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
	OrderId   string    `json:"order_id"`
	OrderNo   string    `json:"order_no"`
	Symbol    string    `json:"symbol"`
	Action    Action    `json:"action"`
	Side      Side      `json:"side"`
	OrderType OrderType `json:"order_type"`
}

type OrderType string

const (
	TypeLimit          OrderType = "TypeLimit"
	TypeLimitIoc       OrderType = "TypeLimitIoc"
	TypeMarket         OrderType = "TypeMarket"
	TypeMarketTop5     OrderType = "TypeMarketTop5"
	TypeMarketTop10    OrderType = "TypeMarketTop10"
	TypeMarketOpponent OrderType = "TypeMarketOpponent"
)

func (a OrderType) String() string {
	return string(a)
}
