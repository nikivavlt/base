package handler

import (
    "crypto/rand"
    "encoding/hex"
    "encoding/json"
    "net/http"
    "strings"

    "github.com/nikivavlt/base/auth/internal/password"
)

// generateRefreshToken creates a cryptographically random opaque token
func generateRefreshToken() (string, error) {
    b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
        return "", err
    }
    return hex.EncodeToString(b), nil // 64 char hex string
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    var body struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        writeError(w, http.StatusBadRequest, "invalid request body")
        return
    }

    body.Email = strings.ToLower(strings.TrimSpace(body.Email))

    // Fetch user
    user, err := h.queries.GetUserByEmail(r.Context(), body.Email)
    if err != nil {
        // Same error for wrong email OR wrong password — prevents user enumeration
        writeError(w, http.StatusUnauthorized, "invalid credentials")
        return
    }

    // Verify password
    if !password.Verify(body.Password, user.Password) {
        writeError(w, http.StatusUnauthorized, "invalid credentials")
        return
    }

    // Sign access token
    accessToken, err := h.jwt.SignAccess(user.ID, user.Email)
    if err != nil {
        writeError(w, http.StatusInternalServerError, "failed to sign token")
        return
    }

    // Generate refresh token
    refreshToken, err := generateRefreshToken()
    if err != nil {
        writeError(w, http.StatusInternalServerError, "failed to generate refresh token")
        return
    }

    // Store in Redis with 7d TTL
    if err := h.redis.Save(r.Context(), refreshToken, user.ID); err != nil {
        writeError(w, http.StatusInternalServerError, "failed to save session")
        return
    }

    // Refresh token → httpOnly cookie
    setRefreshCookie(w, refreshToken)

    // Access token → JSON body (frontend stores in memory)
    writeJSON(w, http.StatusOK, map[string]any{
        "access_token": accessToken,
        "user": map[string]any{
            "id":    user.ID,
            "email": user.Email,
        },
    })
}