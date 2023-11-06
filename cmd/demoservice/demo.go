package main

import (
	"go-demo-svc/mainservice"
	"log"
)

func main() {
	workersNum := 4
	ms, err := mainservice.NewMainService()
	if err != nil {
		log.Println(err)
	}

	ms.Start(workersNum)

}
