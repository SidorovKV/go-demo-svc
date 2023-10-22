package natsclient

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
	"go-demo-svc/cache"
	"go-demo-svc/model"
	"log"
	"sync"
)

const (
	clusterID = "test-cluster"
	clientID  = "test-client"
	channel   = "test-channel"
	durableID = "test-durable"
	queueID   = "test-queue"
)

type PersistenceService interface {
	AddOrder(model.Order) error
	RestoreCache() (map[string]model.Order, error)
}

type NatsClient struct {
	stan.Conn
	validate *validator.Validate
	PersistenceService
}

func NewNatsClient(db PersistenceService) (*NatsClient, error) {
	streamClient, err := stan.Connect(
		clusterID,
		clientID,
		//stan.NatsURL("nats://nats:4222"),
		stan.NatsURL("nats://localhost:4222"),
	)
	if err != nil {
		return nil, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	return &NatsClient{streamClient, validate, db}, nil
}

func (nc *NatsClient) Run(cache cache.CacheService, workersNum int) {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i <= workersNum; i++ {
		wg.Add(1)
		go nc.runWorker(cache, wg, mu, i)
	}
	wg.Wait()
}

func (nc *NatsClient) runWorker(
	cache cache.CacheService,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	workerID int,
) {

	defer wg.Done()

	_, err := nc.QueueSubscribe(channel, queueID, nc.processOrder(cache, mu),
		stan.DurableName(durableID),
		stan.MaxInflight(20),
	)
	if err != nil {
		log.Printf("WorkerID: %v, QueueSubscribe: %v", workerID, err)
		if err := nc.Close(); err != nil {
			log.Printf("WorkerID: %v, conn.Close error: %v", workerID, err)
		}
	}
}

func (nc *NatsClient) processOrder(cache cache.CacheService, mu *sync.Mutex) stan.MsgHandler {
	return func(msg *stan.Msg) {
		var order model.Order
		err := json.Unmarshal(msg.Data, &order)
		if err == nil {
			err = nc.validate.Struct(order)
			if err != nil {
				log.Println("Not valid order received")
			} else {
				mu.Lock()
				cache.Add(order.OrderUID, order)
				mu.Unlock()
				err = nc.AddOrder(order)
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			log.Println(err)
		}
	}
}
