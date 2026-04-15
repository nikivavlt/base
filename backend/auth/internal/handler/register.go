package handler

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/nikivavlt/base/auth/internal/password"
	"github.com/nikivavlt/base/auth/internal/db"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
    var body struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        writeError(w, http.StatusBadRequest, "invalid request body")
        return
    }

    // Validate
    body.Email = strings.ToLower(strings.TrimSpace(body.Email))
    if body.Email == "" || !strings.Contains(body.Email, "@") {
        writeError(w, http.StatusBadRequest, "valid email is required")
        return
    }
    if len(body.Password) < 8 {
        writeError(w, http.StatusBadRequest, "password must be at least 8 characters")
        return
    }

    // Hash password
    hashed, err := password.Hash(body.Password)
    if err != nil {
        writeError(w, http.StatusInternalServerError, "failed to hash password")
        return
    }

    // Insert user
    user, err := h.queries.CreateUser(r.Context(), db.CreateUserParams{
        Email:    body.Email,
        Password: hashed,
    })
    if err != nil {
        // Postgres unique violation code = 23505
        if strings.Contains(err.Error(), "23505") {
            writeError(w, http.StatusConflict, "email already registered")
            return
        }
		println(err.Error())
        writeError(w, http.StatusInternalServerError, "failed to create user")
        return
    }

    writeJSON(w, http.StatusCreated, map[string]any{
        "id":    user.ID,
        "email": user.Email,
    })
}