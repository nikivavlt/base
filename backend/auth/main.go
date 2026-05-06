package main

import (
    "context"
    "net/http"
    "log"
    "os"
    "os/signal"
    "time"
    "syscall"

    "github.com/nikivavlt/base/auth/internal/db"
    "github.com/nikivavlt/base/auth/internal/handler"
    jwtpkg "github.com/nikivavlt/base/auth/internal/jwt"
    "github.com/nikivavlt/base/auth/internal/metrics"
    "github.com/nikivavlt/base/auth/internal/redis"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    dbURL     := os.Getenv("DB_URL")
    redisURL  := os.Getenv("REDIS_URL")
    jwtSecret := os.Getenv("JWT_SECRET")

    if dbURL == "" || redisURL == "" || jwtSecret == "" {
        log.Fatal("DB_URL, REDIS_URL and JWT_SECRET are required")
    }

    database, queries := db.Connect(dbURL)
    defer database.Close()

    redis, err := redis.New(redisURL)
    if err != nil {
        log.Fatalf("redis: %v", err)
    }

    jwt := jwtpkg.NewManager(jwtSecret)
    h   := handler.New(queries, redis, jwt)

    mux := http.NewServeMux()
    mux.Handle("/metrics", promhttp.Handler())
    mux.Handle("/", handler.NewRouter(h))

    srv := &http.Server{
        Addr:         ":8081",
        Handler:      metrics.Middleware(handler.WithCORS(mux)),
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    go func() {
        log.Println("🚀 Auth service running on http://localhost:8081")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("server error: %v", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    srv.Shutdown(ctx)
    log.Println("✅ Auth service stopped")
}