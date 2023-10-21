package model

type Delivery struct {
	Name    string `json:"name" db:"name" fake:"{name}" validate:"required"`
	Phone   string `json:"phone" db:"phone" fake:"+{phone}" validate:"required"`
	Zip     string `json:"zip" db:"zip" fake:"{zip}" validate:"required"`
	City    string `json:"city" db:"city" fake:"{city}" validate:"required"`
	Address string `json:"address" db:"address" fake:"{street} {number:1,300}" validate:"required"`
	Region  string `json:"region" db:"region" fake:"{state}" validate:"required"`
	Email   string `json:"email" db:"email" fake:"{email}" validate:"required,email"`
}
