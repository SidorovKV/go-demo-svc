package dbservice

import (
	"github.com/jackc/pgx/v5"
	"go-demo-svc/model"
)

type RowMapper interface {
	RowsToOrders(pgx.Rows) ([]model.Order, error)
}

type Mapper struct{}

func NewMapper() Mapper {
	return Mapper{}
}

func (Mapper) RowsToOrders(rows pgx.Rows) ([]model.Order, error) {
	return pgx.CollectRows(rows, pgx.RowToStructByName[model.Order])
}
