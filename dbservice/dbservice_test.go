package dbservice

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"go-demo-svc/mocks"
	"go-demo-svc/model"
	"go.uber.org/mock/gomock"
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

func TestDbService_AddOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbmock := mocks.NewMockPgxIface(ctrl)
	mapper := mocks.NewMockRowMapper(ctrl)
	mockedDb := DbService{dbmock, mapper}
	batchResMock := mocks.NewMockBatchResults(ctrl)

	dbmock.EXPECT().SendBatch(context.Background(), gomock.Any()).Return(batchResMock).Times(1)
	batchResMock.EXPECT().Close().Times(1)

	err := mockedDb.AddOrder(o)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDbService_RestoreCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orders := []model.Order{o}
	expect := make(map[string]model.Order)
	expect[o.OrderUID] = o

	dbmock := mocks.NewMockPgxIface(ctrl)
	mockmapper := mocks.NewMockRowMapper(ctrl)
	mockedDb := DbService{dbmock, mockmapper}
	rowsmock := mocks.NewMockRows(ctrl)

	dbmock.EXPECT().Query(context.Background(), gomock.Any()).Return(rowsmock, nil).Times(1)
	mockmapper.EXPECT().RowsToOrders(rowsmock).Return(orders, nil).Times(1)

	out, err := mockedDb.RestoreCache()
	if err != nil {
		t.Fatal(err)
	}

	oout, ok := out[o.OrderUID]

	if ok && reflect.DeepEqual(oout, o) {
		t.Log("cache restored")
	} else {
		t.Fatal("cache not restored")
	}

	errexp := gofakeit.ErrorDatabase()
	dbmock.EXPECT().Query(context.Background(), gomock.Any()).Return(nil, errexp).Times(1)

	_, err = mockedDb.RestoreCache()
	if err == nil {
		t.Fatal("err expected")
	} else if !errors.Is(errexp, err) {
		t.Fatal("not expected error")
	} else {
		t.Log("ok")
	}

	errexp = gofakeit.Error()
	dbmock.EXPECT().Query(context.Background(), gomock.Any()).Return(rowsmock, nil).Times(1)
	mockmapper.EXPECT().RowsToOrders(rowsmock).Return(nil, errexp).Times(1)

	_, err = mockedDb.RestoreCache()
	if err == nil {
		t.Fatal("err expected")
	} else if !errors.Is(errexp, err) {
		t.Fatal("not expected error")
	} else {
		t.Log("ok")
	}
}
