package model

type Item struct {
	ChrtID      int    `json:"chrt_id" db:"chrt_id" fake:"{uint64:1,9223372036854775806}" validate:"required"`
	TrackNumber string `json:"track_number" db:"track_number" validate:"required"`
	Price       int    `json:"price" db:"price" fake:"{uint16}" validate:"required"`
	Rid         string `json:"rid" db:"rid" fake:"{uuid}" validate:"required"`
	Name        string `json:"name" db:"name" fake:"{minecraftarmorpart}" validate:"required"`
	Sale        int    `json:"sale" db:"sale" validate:"required"`
	Size        string `json:"size" db:"size" fake:"XXL" validate:"required"`
	TotalPrice  int    `json:"total_price,omitempty" db:"-"`
	NmID        int    `json:"nm_id" db:"nm_id" fake:"{uint16}" validate:"required"`
	Brand       string `json:"brand" db:"brand" fake:"{beername}" validate:"required"`
	Status      int    `json:"status" db:"status" fake:"{uint8}" validate:"required"`
}
