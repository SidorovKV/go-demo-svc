package cache

import (
	"go-demo-svc/model"
)

type CacheService interface {
	Get(orderUID string) (model.Order, bool)
	Add(OrderUID string, order model.Order)
}

type Cache struct {
	orders map[string]model.Order
}

func NewCache(orders map[string]model.Order) *Cache {
	return &Cache{
		orders,
	}
}

func (cs *Cache) Get(orderUID string) (model.Order, bool) {
	order, ok := cs.orders[orderUID]
	return order, ok
}

func (cs *Cache) Add(OrderUID string, order model.Order) {
	cs.orders[OrderUID] = order
}
