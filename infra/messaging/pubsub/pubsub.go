package pubsubmessaging

import (
	"context"
	"fmt"

	"calculator.com/infra/messaging"
	"cloud.google.com/go/pubsub"
)

type pubsubMessaging struct {
	client *pubsub.Client
}

var _ messaging.Messaging = (*pubsubMessaging)(nil)

func (p pubsubMessaging) Publish(ctx context.Context, topic string, msg interface{}) {
	topicName := p.client.Topic(topic)
	res := topicName.Publish(ctx, &pubsub.Message{
		Data: []byte(msg.(string)),
	})

	_, err := res.Get(ctx)
	if err != nil {
		fmt.Printf("Error publishing message: %v", err)
	}

	fmt.Printf("Published message: %v topic: %v", msg, topic)
}

func (p pubsubMessaging) Subscribe(ctx context.Context, topic string, handler messaging.MessageHandler) {
	sub := p.client.Subscription(topic)
	sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		handler(ctx, string(msg.Data))
		msg.Ack()
	})
}

func NewPubsubMessaging(client *pubsub.Client) *pubsubMessaging {
	return &pubsubMessaging{
		client: client,
	}
}
