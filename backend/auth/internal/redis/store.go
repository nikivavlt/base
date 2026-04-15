package redis

import (
    "context"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

const refreshTokenTTL = 7 * 24 * time.Hour

type Store struct {
    client *redis.Client
}

func New(redisURL string) (*Store, error) {
    opts, err := redis.ParseURL(redisURL)
    if err != nil {
        return nil, fmt.Errorf("parsing redis URL: %w", err)
    }

    client := redis.NewClient(opts)

    // Verify connection
    if err := client.Ping(context.Background()).Err(); err != nil {
        return nil, fmt.Errorf("connecting to redis: %w", err)
    }

    fmt.Println("✅ Connected to Redis")
    return &Store{client: client}, nil
}

// key format: refresh:<token>
func tokenKey(token string) string {
    return "refresh:" + token
}

// Save stores userID against the refresh token with TTL
func (s *Store) Save(ctx context.Context, token string, userID int64) error {
    return s.client.Set(ctx,
        tokenKey(token),
        userID,
        refreshTokenTTL,
    ).Err()
}

// Get returns the userID for a refresh token
func (s *Store) Get(ctx context.Context, token string) (int64, error) {
    val, err := s.client.Get(ctx, tokenKey(token)).Int64()
    if err == redis.Nil {
        return 0, fmt.Errorf("refresh token not found or expired")
    }
    return val, err
}

// Delete revokes a refresh token (logout)
func (s *Store) Delete(ctx context.Context, token string) error {
    return s.client.Del(ctx, tokenKey(token)).Err()
}