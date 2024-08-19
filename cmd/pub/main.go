package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/manochatt/line-noti/domain"
	"github.com/manochatt/line-noti/utils/nats_utils"
)

func main() {
	jsCtx, err := nats_utils.InitialNatServer()
	if err != nil {
		log.Fatal("Error", err)
	}

	demo := domain.LineMessageDTO{
		ToID:      "Ub0d85d97a8a3688c45662a2241d313e9",
		ProjectID: "66b47943bef4dd43c5a1b7e6",
		MessageValue: domain.MessageValue{
			Title:         "Test Coffee",
			Place:         "BKK",
			StartDateTime: "10.00",
			EndDateTime:   "12:00",
		},
	}

	demoData, err := json.Marshal(demo)
	if err != nil {
		log.Fatal("Error")
	}

	_, err = jsCtx.Publish("demo.data", []byte(demoData))
	if err != nil {
		fmt.Println("Can not publish a new message to NAT server")
	}
	fmt.Println("Done")
}
