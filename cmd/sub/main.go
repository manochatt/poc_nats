package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/manochatt/line-noti/utils/nats_utils"
	"github.com/nats-io/nats.go"
)

func main() {
	ctx := context.Background()

	jsCtx, err := nats_utils.InitialNatServer()
	if err != nil {
		log.Fatal("Error", err)
	}

	_, err = createConsumer(ctx, jsCtx, "demo_group", "demo")
	if err != nil {
		log.Fatal("Error", err)
	}

	defer func() {
		if err := deleteConsumer(ctx, jsCtx, "demo_group", "demo"); err != nil {
			log.Fatal("Error", err)
		}
	}()

	pullSub, err := subscribe(ctx, jsCtx, "demo.*", "demo_group", "demo")
	if err != nil {
		log.Fatal("Error", err)
	}

	msgCh := make(chan *nats.Msg)

	go func() {
		for {
			fetchCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			msg, err := fetchOne(fetchCtx, pullSub)
			cancel()
			if err != nil {
				log.Printf("No new message in recent 5 seconds: %s", err)
				time.Sleep(1 * time.Second)
				continue
			}

			msgCh <- msg
		}
	}()

	for msg := range msgCh {
		fmt.Println("✅", string(msg.Data))

		var messageObj map[string]interface{}
		err := json.Unmarshal(msg.Data, &messageObj)
		if err != nil {
			log.Fatal("Error cannot unmarshal messages:", err)
		}

		if messageObj["Successful"] != true {
			continue
		}

		// Acknowledge the message manually
		if err := msg.Ack(); err != nil {
			log.Printf("Error acknowledging message: %v", err)
		}
	}
}

func createConsumer(ctx context.Context, jsCtx nats.JetStreamContext, consumerGroupName, streamName string) (*nats.ConsumerInfo, error) {
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

func deleteConsumer(ctx context.Context, jsCtx nats.JetStreamContext, consumerGroupName, streamName string) error {
	err := jsCtx.DeleteConsumer(streamName, consumerGroupName, nats.Context(ctx))
	if err != nil {
		return fmt.Errorf("delete consumer: %w", err)
	}

	return nil
}

func subscribe(ctx context.Context, js nats.JetStreamContext, subject, consumerGroupName, streamName string) (*nats.Subscription, error) {
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

func fetchOne(ctx context.Context, pullSub *nats.Subscription) (*nats.Msg, error) {
	msgs, err := pullSub.Fetch(1, nats.Context(ctx))
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}
	if len(msgs) == 0 {
		return nil, errors.New("no messages")
	}

	return msgs[0], nil
}
