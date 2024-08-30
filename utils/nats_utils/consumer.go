package nats_utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	line_models "github.com/manochatt/line-noti/domain/line/models"
	line_requests "github.com/manochatt/line-noti/domain/line/requests"
	"github.com/manochatt/line-noti/modules/line/repository"
	"github.com/manochatt/line-noti/modules/line/usecase"
	"github.com/manochatt/line-noti/mongo"
	"github.com/nats-io/nats.go"
)

func Consumer(timeout time.Duration, db mongo.Database) {
	ctx := context.Background()

	jsCtx, err := InitialNatServer()
	if err != nil {
		log.Fatal("Error", err)
	}

	_, err = CreateConsumer(ctx, jsCtx, "demo_group", "demo")
	if err != nil {
		log.Fatal("Error", err)
	}

	// Remove consumer before terminate application
	defer func() {
		if err := DeleteConsumer(ctx, jsCtx, "demo_group", "demo"); err != nil {
			log.Fatal("Error", err)
		}
	}()

	pullSub, err := Subscribe(ctx, jsCtx, "demo.*", "demo_group", "demo")
	if err != nil {
		log.Fatal("Error", err)
	}

	msgCh := make(chan *nats.Msg)

	go func() {
		for {
			fetchCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
			msg, err := FetchOne(fetchCtx, pullSub)
			cancel()
			if err != nil {
				log.Printf("⌛️ No new message in recent 5 minutes: %s", err)
				time.Sleep(1 * time.Second)
				continue
			}

			msgCh <- msg
		}
	}()

	lr := repository.NewLineRepository(db, line_models.CollectionLineTemplate)
	lu := usecase.NewLineUsecase(lr, timeout)

	for msg := range msgCh {
		fmt.Println("✅", string(msg.Data))

		var notificationData line_requests.LineMessageRequest
		err := json.Unmarshal(msg.Data, &notificationData)
		if err != nil {
			fmt.Printf("Error cannot unmarshal dto: %v", err)
			continue
		}
		err = notificationData.Validate()
		if err != nil {
			fmt.Printf("Error validation error: %v", err)
			continue
		}

		err = lu.SendMessage(ctx, notificationData)
		if err != nil {
			fmt.Printf("Cannot send notification message: %v", err)
			continue
		}

		// Acknowledge the message manually
		if err := msg.Ack(); err != nil {
			log.Printf("Error acknowledging message: %v", err)
		}
	}
}

func CreateConsumer(ctx context.Context, jsCtx nats.JetStreamContext, consumerGroupName, streamName string) (*nats.ConsumerInfo, error) {
	consumer, err := jsCtx.AddConsumer(streamName, &nats.ConsumerConfig{
		Durable:       consumerGroupName,      // durable name is the same as consumer group name
		DeliverPolicy: nats.DeliverAllPolicy,  // deliver all messages, even if they were sent before the consumer was created
		AckPolicy:     nats.AckExplicitPolicy, // ack messages manually
		AckWait:       5 * time.Second,        // wait for ack for 5 seconds
		MaxAckPending: -1,                     // unlimited number of pending acks
		MaxDeliver:    3,                      // maximum deliver attempts
	}, nats.Context(ctx))
	if err != nil {
		return nil, fmt.Errorf("add consumer: %w", err)
	}

	return consumer, nil
}

func DeleteConsumer(ctx context.Context, jsCtx nats.JetStreamContext, consumerGroupName, streamName string) error {
	err := jsCtx.DeleteConsumer(streamName, consumerGroupName, nats.Context(ctx))
	if err != nil {
		return fmt.Errorf("delete consumer: %w", err)
	}

	return nil
}

func Subscribe(ctx context.Context, js nats.JetStreamContext, subject, consumerGroupName, streamName string) (*nats.Subscription, error) {
	pullSub, err := js.PullSubscribe(
		subject,
		consumerGroupName,
		nats.ManualAck(),                         // ack messages manually
		nats.Bind(streamName, consumerGroupName), // bind consumer to the stream
		nats.Context(ctx),                        // use context to cancel the subscription
	)
	if err != nil {
		return nil, fmt.Errorf("pull subscribe: %w", err)
	}

	return pullSub, nil
}

func FetchOne(ctx context.Context, pullSub *nats.Subscription) (*nats.Msg, error) {
	msgs, err := pullSub.Fetch(1, nats.Context(ctx))
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}
	if len(msgs) == 0 {
		return nil, errors.New("no messages")
	}

	return msgs[0], nil
}
