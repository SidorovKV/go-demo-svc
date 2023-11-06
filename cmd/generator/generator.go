package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"go-demo-svc/model"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	clusterID = "test-cluster"
	clientID  = "test-pusher"
	channel   = "test-channel"
)

type trash struct {
	price       float64
	sale        int
	name        string
	despription string
	created     time.Time
}

func main() {
	var order model.Order
	var trash trash
	uuids := make([]string, 0, 10000)

	counter := 0
	validate := validator.New(validator.WithRequiredStructEnabled())

	streamClient, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(stan.DefaultNatsURL),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer streamClient.Close()
	for i := 0; i < 10000; i++ {
		t := rand.Intn(10)
		if t > 5 {
			gofakeit.Struct(&trash)
			out, err := json.Marshal(trash)
			if err != nil {
				log.Fatal(err)
			}
			err = streamClient.Publish(channel, out)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			gofakeit.Struct(&order)
			order.Payment.Transaction = order.OrderUID
			gt := 0
			for n, _ := range order.Items {
				sale := rand.Intn(65)
				order.Items[n].TrackNumber = order.TrackNumber
				order.Items[n].Sale = sale
				order.Items[n].TotalPrice = (order.Items[n].Price * sale) / 100
				gt += order.Items[n].TotalPrice
			}
			order.Payment.GoodsTotal = gt
			out, err := json.Marshal(order)
			if err != nil {
				log.Fatal(err)
			}
			err = streamClient.Publish(channel, out)
			if err != nil {
				log.Fatal(err)
			}

			if validate.Struct(order) == nil {
				uuids = append(uuids, order.OrderUID)
				counter++
			}
		}
	}
	builder := strings.Builder{}
	builder.WriteString(
		`math.randomseed(os.time())
request = function() 
   
   local r = {
`)
	for n, uuid := range uuids {
		builder.WriteString(fmt.Sprintf(`{uuid="%s"}`, uuid))
		if n < len(uuids)-1 {
			builder.WriteString(",")
		}
	}
	builder.WriteString(
		`}
	url_path = "/api/order/?uid=" .. r[math.random(1,#r)].uuid
	return wrk.format("GET", url_path)
end
`)
	f, _ := os.Create("script.lua")
	defer f.Close()
	f.WriteString(builder.String())
	fmt.Println(counter)
}
