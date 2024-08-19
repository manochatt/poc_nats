package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/manochatt/line-noti/utils/nats_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Demo struct {
	ID         primitive.ObjectID
	Message    string
	Successful bool
}

func main() {
	jsCtx, err := nats_utils.InitialNatServer()
	if err != nil {
		log.Fatal("Error", err)
	}

	demo := Demo{
		ID:         primitive.NewObjectID(),
		Message:    "test",
		Successful: false,
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
