package handler

import "net/http"

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
    refreshToken, err := getRefreshCookie(r)
    if err != nil {
        // Already logged out — still return 200
        w.WriteHeader(http.StatusOK)
        return
    }

    // Revoke from Redis — token is dead immediately
    _ = h.redis.Delete(r.Context(), refreshToken)

    clearRefreshCookie(w)
    w.WriteHeader(http.StatusOK)
}