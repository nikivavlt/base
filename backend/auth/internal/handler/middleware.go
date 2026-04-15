package handler

import (
    "net/http"
    "os"
)

// Allow-Credentials: true is required — without it, browsers refuse to send cookies cross-origin.
func WithCORS(next http.Handler) http.Handler {
    origin := os.Getenv("CORS_ORIGIN")
    if origin == "" {
        origin = "http://localhost:5173"
    }

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Allow-Credentials", "true") // required for cookies

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }

        next.ServeHTTP(w, r)
    })
}