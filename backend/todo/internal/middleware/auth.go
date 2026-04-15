package middleware

import (
    "context"
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

// Context key type — avoids collisions with other context values
type contextKey string

const ClaimsKey contextKey = "claims"

type Claims struct {
    UserID int64  `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

type AuthMiddleware struct {
    secret []byte
}

func NewAuth(secret string) *AuthMiddleware {
    return &AuthMiddleware{secret: []byte(secret)}
}

func (a *AuthMiddleware) Protect(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract token from Authorization: Bearer <token>
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            writeError(w, http.StatusUnauthorized, "missing authorization header")
            return
        }

        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            writeError(w, http.StatusUnauthorized, "invalid authorization format")
            return
        }

        tokenStr := parts[1]

        // Parse and validate
        claims := &Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims,
            func(t *jwt.Token) (any, error) {
                if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("unexpected signing method")
                }
                return a.secret, nil
            },
        )
        if err != nil || !token.Valid {
            writeError(w, http.StatusUnauthorized, "invalid or expired token")
            return
        }

        // Check expiry explicitly
        if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
            writeError(w, http.StatusUnauthorized, "token expired")
            return
        }

        // Attach claims to request context
        ctx := context.WithValue(r.Context(), ClaimsKey, claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// GetClaims extracts claims from context — used in handlers
func GetClaims(r *http.Request) *Claims {
    claims, _ := r.Context().Value(ClaimsKey).(*Claims)
    return claims
}

// helper — keeps middleware self-contained
func writeError(w http.ResponseWriter, status int, msg string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    fmt.Fprintf(w, `{"error":%q}`, msg)
}