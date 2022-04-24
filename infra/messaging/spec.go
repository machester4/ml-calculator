package messaging

import "context"

type Messaging interface {
	Publish(ctx context.Context, topic string, msg interface{})
	Subscribe(ctx context.Context, topic string, handler MessageHandler)
}

type MessageHandler func(ctx context.Context, msg interface{})
