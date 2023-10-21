package model

type Order struct {
	OrderUID          string   `json:"order_uid" db:"order_uid" fake:"{uuid}" validate:"required"`
	TrackNumber       string   `json:"track_number" db:"track_number" fake:"{isin}" validate:"required"`
	Entry             string   `json:"entry" db:"entry" validate:"required"`
	Delivery          Delivery `json:"delivery" db:"delivery" validate:"required"`
	Payment           Payment  `json:"payment" db:"payment" validate:"required"`
	Items             []Item   `json:"items" db:"items" fakesize:"1,6" validate:"required,dive,required"`
	Locale            string   `json:"locale" db:"locale" fake:"{languageabbreviation}" validate:"required"`
	InternalSignature string   `json:"internal_signature" db:"internal_signature"`
	CustomerID        string   `json:"customer_id" db:"customer_id" validate:"required"`
	DeliveryService   string   `json:"delivery_service" db:"delivery_service" validate:"required"`
	Shardkey          string   `json:"shardkey" db:"shardkey" fake:"{number:1,12}" validate:"required"`
	SmID              int      `json:"sm_id" db:"sm_id" fake:"{uint8}" validate:"required"`
	DateCreated       string   `json:"date_created" db:"date_created" validate:"required" fake:"{year}-{month}-{day}T{hour}:{minute}:{second}Z"`
	OofShard          string   `json:"oof_shard" db:"oof_shard" fake:"{number:1,12}" validate:"required"`
}
