package handler

import (
    "encoding/json"
    "net/http"

    "github.com/nikivavlt/base/auth/internal/db"
    jwtpkg "github.com/nikivavlt/base/auth/internal/jwt"
    "github.com/nikivavlt/base/auth/internal/redis"
)

type Handler struct {
    queries *db.Queries
    redis   *redis.Store
    jwt     *jwtpkg.Manager
}

func New(queries *db.Queries, redis *redis.Store, jwt *jwtpkg.Manager) *Handler {
    return &Handler{queries: queries, redis: redis, jwt: jwt}
}

// ── helpers ────────────────────────────────────────────────────────────────

func writeJSON(w http.ResponseWriter, status int, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
    writeJSON(w, status, map[string]string{"error": msg})
}