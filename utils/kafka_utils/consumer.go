package kafka_utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/manochatt/line-noti/bootstrap"
	"github.com/manochatt/line-noti/config"
	"github.com/manochatt/line-noti/domain/models"
	"github.com/manochatt/line-noti/mongo"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TopicSubscribe(env *bootstrap.Env, db mongo.Database) {
	cfg := config.KafkaConnCfg{
		Url:   env.KafkaURL,
		Topic: "shop",
	}
	conn := KafkaConn(cfg)
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal("fail to close connection:", err)
		}
	}()

	lineTemplateCh := make(chan models.LineTemplate)

	go ReadMessage(conn, lineTemplateCh)
	LineTemplateHandler(db, lineTemplateCh)

	close(lineTemplateCh)

}

func ReadMessage(conn *kafka.Conn, lineTemplateCh chan models.LineTemplate) {
	var lineTemplate models.LineTemplate
	for {
		message, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}

		err = json.Unmarshal(message.Value, &lineTemplate)
		if err != nil {
			log.Fatal("Error unmarshaling message:", err)
		}

		lineTemplateCh <- lineTemplate
		fmt.Println("✅", string(message.Value))
	}
}

func LineTemplateHandler(db mongo.Database, lineTemplateCh chan models.LineTemplate) {
	collection := db.Collection(models.CollectionLineTemplate)
	c := context.Background()
	for lc := range lineTemplateCh {
		time.Sleep(time.Millisecond)
		var lineTemplates []models.LineTemplate
		fmt.Println("========================================================================")

		cursor, err := collection.Find(c, bson.M{"_id": lc.ID})
		if err != nil {
			log.Fatal("Error cannot find line template", err)
		}

		cursor.All(c, &lineTemplates)
		if lineTemplates != nil {
			continue
		}

		lc.ID = primitive.NewObjectID()

		_, err = collection.InsertOne(c, &lc)
		if err != nil {
			log.Fatal("Error cannot create line template", err)
		}
	}
}
