package main

import (
    "log"
    "net/http"
    "os"

    "github.com/nikivavlt/base/todo/internal/db"
    "github.com/nikivavlt/base/todo/internal/handler"
    "github.com/nikivavlt/base/todo/internal/kafka"
    "github.com/nikivavlt/base/todo/internal/middleware"
)

func main() {
    dbURL      := os.Getenv("DB_URL")
    jwtSecret  := os.Getenv("JWT_SECRET")
    kafkaBroker := os.Getenv("KAFKA_BROKER")

    if dbURL == "" {
        log.Fatal("DB_URL environment variable is required")
    }
    if kafkaBroker == "" {
        kafkaBroker = "kafka:9092"
    }

    database, queries := db.Connect(dbURL)
    defer database.Close()

    producer := kafka.NewProducer(kafkaBroker)
    defer producer.Close()

    auth   := middleware.NewAuth(jwtSecret)
    h      := handler.New(queries, producer)
    router := handler.NewRouter(h, auth)

    srv := &http.Server{
        Addr:    ":8080",
        Handler: handler.WithCORS(router),
    }

    log.Println("🚀 Server running on http://localhost:8080")
    if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}