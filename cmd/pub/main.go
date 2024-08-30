package main

import (
	"encoding/json"
	"fmt"
	"log"

	line_requests "github.com/manochatt/line-noti/domain/line/requests"
	"github.com/manochatt/line-noti/utils/nats_utils"
)

func main() {
	jsCtx, err := nats_utils.InitialNatServer()
	if err != nil {
		log.Fatal("Error", err)
	}

	demo := line_requests.LineMessageRequest{
		ToID:      "Ub0d85d97a8a3688c45662a2241d313e9",
		ProjectID: "66b47943bef4dd43c5a1b7e6",
		MessageValue: line_requests.MessageValue{
			Title:         "Test Coffee",
			Place:         "BKK",
			StartDateTime: "2024-07-17T05:00:00+07:00",
			EndDateTime:   "2024-07-17T06:00:00+07:00",
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
