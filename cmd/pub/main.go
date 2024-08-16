package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/manochatt/line-noti/demo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Demo struct {
	ID         primitive.ObjectID
	Message    string
	SuccessFul bool
}

func main() {
	jsCtx, err := demo.InitialNatServer()
	if err != nil {
		log.Fatal("Error", err)
	}

	demo := Demo{
		ID:         primitive.NewObjectID(),
		Message:    "test",
		SuccessFul: true,
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
