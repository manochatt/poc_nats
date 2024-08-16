package main

import (
	"log"
	"time"

	"github.com/manochatt/line-noti/config"
	"github.com/manochatt/line-noti/domain"
	"github.com/manochatt/line-noti/utils"
	"github.com/manochatt/line-noti/utils/kafka_utils"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Connection part
	cfg := config.KafkaConnCfg{
		Url:   "host.docker.internal:9092",
		Topic: "shop",
	}
	conn := kafka_utils.KafkaConn(cfg)

	// Check topic if already exists or not
	if !kafka_utils.IsTopicAlreadyExists(conn, cfg.Topic) {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             cfg.Topic,
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}

	// Mock data
	data := func() []kafka.Message {
		products := []domain.LineTemplate{
			{
				ToID: primitive.NewObjectID(),
				Messages: []domain.Message{
					{
						Type:    "A",
						AltText: "B",
					},
				},
			},
			{
				ToID: primitive.NewObjectID(),
				Messages: []domain.Message{
					{
						Type:    "A",
						AltText: "C",
					},
				},
			},
			{
				ToID: primitive.NewObjectID(),
				Messages: []domain.Message{
					{
						Type:    "A",
						AltText: "D",
					},
				},
			},
		}

		// Convert into kafka.Message{}
		messages := make([]kafka.Message, 0)
		for _, p := range products {
			messages = append(messages, kafka.Message{
				Value: utils.CompressToJsonBytes(&p),
			})
		}
		return messages
	}()

	// Set timeout
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(data...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	// Close connection
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
