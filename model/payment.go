package model

type Payment struct {
	Transaction  string `json:"transaction" db:"transaction" validate:"required"`
	RequestID    string `json:"request_id" db:"request_id" fake:"{number}"`
	Currency     string `json:"currency" db:"currency" fake:"{currencyshort}" validate:"required"`
	Provider     string `json:"provider" db:"provider" validate:"required"`
	Amount       int    `json:"amount" db:"-" fake:"{uint16}"`
	PaymentDt    int    `json:"payment_dt" db:"payment_dt" fake:"{uint32}" validate:"required"`
	Bank         string `json:"bank" db:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost" fake:"{uint16}" validate:"required"`
	GoodsTotal   int    `json:"goods_total" db:"-" fake:"{uint16}"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee" fake:"{uint16}" validate:"required"`
}
