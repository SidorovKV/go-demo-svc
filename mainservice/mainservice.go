package mainservice

import (
	"encoding/json"
	"go-demo-svc/cache"
	"go-demo-svc/dbservice"
	"go-demo-svc/natsclient"
	"log"
	"net/http"
)

type MainService struct {
	cache cache.CacheService
	nc    *natsclient.NatsClient
}

func NewMainService() (*MainService, error) {
	db, err := dbservice.NewDbService()
	if err != nil {
		return nil, err
	}

	orders, err := db.RestoreCache()
	if err != nil {
		return nil, err
	}

	ca := cache.NewCache(orders)

	nc, err := natsclient.NewNatsClient(db)
	if err != nil {
		return nil, err
	}
	out := &MainService{
		ca,
		nc,
	}
	return out, nil
}

func (ms MainService) Start(workersNum int) {
	go func() {
		ms.nc.Run(ms.cache, workersNum)
	}()
	http.HandleFunc("/api/order/", ms.handler)

	if http.ListenAndServe(":8080", nil) != nil {
		log.Fatal("Error starting server")
		return
	}
}

func (ms MainService) handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("uid") {
		uid := r.URL.Query().Get("uid")
		order, ok := ms.cache.Get(uid)
		if ok {
			out, err := json.Marshal(order)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, err := w.Write([]byte(err.Error()))
				if err != nil {
					log.Println(err)
				}
				return
			}
			_, err = w.Write(out)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("Not found"))
			if err != nil {
				log.Println(err)
			}
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Bad request"))
		if err != nil {
			log.Println(err)
		}
		return
	}
}
