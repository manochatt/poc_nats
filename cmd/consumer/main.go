package main

import (
	"fmt"
	"log"

	"github.com/manochatt/line-noti/config"
	"github.com/manochatt/line-noti/utils/kafka_utils"
)

func main() {
	cfg := config.KafkaConnCfg{
		Url:   "host.docker.internal:9092",
		Topic: "shop",
	}
	conn := kafka_utils.KafkaConn(cfg)

	for {
		message, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		fmt.Println(string(message.Value))
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
