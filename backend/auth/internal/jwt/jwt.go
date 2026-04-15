package jwt

import (
    "fmt"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

const accessTokenTTL = 15 * time.Minute

type Claims struct {
    UserID int64  `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

type Manager struct {
    secret []byte
}

func NewManager(secret string) *Manager {
    return &Manager{secret: []byte(secret)}
}

// SignAccess creates a 15min access token
func (m *Manager) SignAccess(userID int64, email string) (string, error) {
    claims := Claims{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed, err := token.SignedString(m.secret)
    if err != nil {
        return "", fmt.Errorf("signing access token: %w", err)
    }
    return signed, nil
}

// Verify parses and validates a token, returns claims if valid
func (m *Manager) Verify(tokenStr string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &Claims{},
        func(t *jwt.Token) (any, error) {
            if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
            }
            return m.secret, nil
        },
    )
    if err != nil {
        return nil, fmt.Errorf("parsing token: %w", err)
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token claims")
    }

    return claims, nil
}