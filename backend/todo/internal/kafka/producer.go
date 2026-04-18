package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

const TopicTodoCreated = "todo.created"

type TodoCreatedEvent struct {
	TodoID int64  `json:"todo_id"`
	Title  string `json:"title"`
}

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(broker string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(broker),
			Topic:                  TopicTodoCreated,
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		},
	}
}

func (p *Producer) PublishTodoCreated(ctx context.Context, event TodoCreatedEvent) {
	payload, err := json.Marshal(event)
	if err != nil {
		log.Printf("kafka: marshal failed: %v", err)
		return
	}

	if err := p.writer.WriteMessages(ctx, kafka.Message{Value: payload}); err != nil {
		log.Printf("kafka: publish failed: %v", err)
	}
}

func (p *Producer) Close() {
	p.writer.Close()
}
