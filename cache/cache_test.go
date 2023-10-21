package cache

import (
	"github.com/brianvoe/gofakeit/v6"
	"go-demo-svc/model"
	"reflect"
	"testing"
)

var (
	d = model.Delivery{
		Name:    "Hipolito Corwin",
		Phone:   "+6552369296",
		Zip:     "14034",
		City:    "Fremont",
		Address: "Trafficway 61",
		Region:  "Indiana",
		Email:   "cristalabbott@bergstrom.net",
	}

	p = model.Payment{
		Transaction:  "90f3eb2b-af41-49af-9e4b-a8e5bd809311",
		RequestID:    "6599238",
		Currency:     "USD",
		Provider:     "bxzxoNAre",
		Amount:       53256,
		PaymentDt:    845790043,
		Bank:         "lItqGB",
		DeliveryCost: 14367,
		GoodsTotal:   32449,
		CustomFee:    6440,
	}

	items = []model.Item{model.Item{
		ChrtID:      324648622006817776,
		TrackNumber: "RUZPUXCVDY36",
		Price:       40000,
		Rid:         "59b89efd-e3c1-4a35-ba9d-970d687106ff",
		Name:        "helmet",
		Sale:        19,
		Size:        "XXL",
		TotalPrice:  32449,
		NmID:        5,
		Brand:       "Racer",
		Status:      207,
	}}

	o = model.Order{
		OrderUID:          "90f3eb2b-af41-49af-9e4b-a8e5bd809311",
		TrackNumber:       "RUZPUXCVDY36",
		Entry:             "vgQY",
		Delivery:          d,
		Payment:           p,
		Items:             items,
		Locale:            "es",
		InternalSignature: "MqxhMmYv",
		CustomerID:        "sAOPj",
		DeliveryService:   "ePUEDSVJSO",
		Shardkey:          "8",
		SmID:              8,
		DateCreated:       "1914-8-31T21:58:6Z",
		OofShard:          "1",
	}
)

func TestCache_Add(t *testing.T) {

	var o2, o3 model.Order
	gofakeit.Struct(&o2)
	gofakeit.Struct(&o3)

	type fields struct {
		orders map[string]model.Order
	}
	type args struct {
		OrderUID string
		order    model.Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"test1", fields{make(map[string]model.Order)}, args{o.OrderUID, o}},
		{"test2", fields{make(map[string]model.Order)}, args{o2.OrderUID, o2}},
		{"test3", fields{make(map[string]model.Order)}, args{o3.OrderUID, o3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &Cache{
				orders: tt.fields.orders,
			}
			cs.Add(tt.args.OrderUID, tt.args.order)
			order, ok := cs.orders[tt.args.OrderUID]
			if !ok {
				t.Errorf("Add() failed to add order to cache")
			}
			if !reflect.DeepEqual(order, tt.args.order) {
				t.Errorf("Add() failed to add order to cache")
			}
			_, ok = cs.orders["trash"]
			if ok {
				t.Errorf("Trash in cache")
			}
		})
	}
}

func TestCache_Get(t *testing.T) {
	orders := make(map[string]model.Order)
	var o2, o3 model.Order
	gofakeit.Struct(&o2)
	gofakeit.Struct(&o3)

	orders[o.OrderUID] = o
	orders[o2.OrderUID] = o2
	orders[o3.OrderUID] = o3

	type fields struct {
		orders map[string]model.Order
	}
	type args struct {
		orderUID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   model.Order
		want1  bool
	}{
		{"test1", fields{orders}, args{o.OrderUID}, o, true},
		{"test2", fields{orders}, args{"trash"}, model.Order{}, false},
		{"test3", fields{orders}, args{o2.OrderUID}, o2, true},
		{"test4", fields{orders}, args{o3.OrderUID}, o3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &Cache{
				orders: tt.fields.orders,
			}
			got, got1 := cs.Get(tt.args.orderUID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewCache(t *testing.T) {
	orders := make(map[string]model.Order)
	orders1 := make(map[string]model.Order)
	orders2 := make(map[string]model.Order)

	var o2, o3 model.Order
	gofakeit.Struct(&o2)
	gofakeit.Struct(&o3)

	orders[o.OrderUID] = o
	orders[o2.OrderUID] = o2
	orders[o3.OrderUID] = o3

	orders1[o.OrderUID] = o
	orders1[o2.OrderUID] = o2

	orders2[o3.OrderUID] = o3

	type args struct {
		orders map[string]model.Order
	}
	tests := []struct {
		name string
		args args
		want *Cache
	}{
		{"test1", args{orders}, &Cache{orders}},
		{"test2", args{orders1}, &Cache{orders1}},
		{"test3", args{orders2}, &Cache{orders2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCache(tt.args.orders); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
