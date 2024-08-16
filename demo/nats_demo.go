package demo

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func InitialNatServer() (nats.JetStreamContext, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	n, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	jsCtx, err := n.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}

	s, err := jsCtx.StreamInfo("demo")
	if s == nil {
		// Create new stream
		_, err := createStream(ctx, jsCtx)
		if err != nil {
			return nil, err
		}
	}

	return jsCtx, nil
}

func createStream(ctx context.Context, jsCtx nats.JetStreamContext) (*nats.StreamInfo, error) {
	stream, err := jsCtx.AddStream(&nats.StreamConfig{
		Name:              "demo",
		Subjects:          []string{"demo.*", "subject.2", "subject.N"},
		Retention:         nats.InterestPolicy, // remove acked messages
		Discard:           nats.DiscardOld,     // when the stream is full, discard old messages
		MaxAge:            7 * 24 * time.Hour,  // max age of stored messages is 7 days
		Storage:           nats.FileStorage,    // type of message storage
		MaxMsgsPerSubject: 100_000_000,         // max stored messages per subject
		MaxMsgSize:        4 << 20,             // max single message size is 4 MB
		NoAck:             false,               // we need the "ack" system for the message queue system
	}, nats.Context(ctx))
	if err != nil {
		return nil, fmt.Errorf("add stream: %w", err)
	}

	return stream, nil
}

func publishMsg(nc *nats.Conn, subject string, payload []byte) error {
	err := nc.Publish(subject, payload)
	if err != nil {
		return fmt.Errorf("publish: %w", err)
	}

	return nil
}
