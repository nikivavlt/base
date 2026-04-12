package main

import (
    "log"
    "net/http"
    "os"

    "github.com/nikivavlt/base/internal/db"
    "github.com/nikivavlt/base/internal/handler"
)

func main() {
    dbURL := os.Getenv("DB_URL")
    if dbURL == "" {
        log.Fatal("DB_URL environment variable is required")
    }

    database, queries := db.Connect(dbURL)
    defer database.Close()

    h := handler.New(queries)
    router := handler.NewRouter(h)

    srv := &http.Server{
        Addr:    ":8080",
        Handler: handler.WithCORS(router),
    }

    log.Println("🚀 Server running on http://localhost:8080")
    if err := srv.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}