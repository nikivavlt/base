package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type TodoCreatedEvent struct {
	TodoID int64  `json:"todo_id"`
	Title  string `json:"title"`
}

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "kafka:9092"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   "todo.created",
		GroupID: "notification-service",
	})
	defer r.Close()

	log.Println("🔔 Notification service listening on topic: todo.created")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("read error: %v", err)
			continue
		}

		var event TodoCreatedEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("unmarshal error: %v", err)
			continue
		}

		log.Printf("📬 New todo created — id: %d, title: %q", event.TodoID, event.Title)
	}
}
