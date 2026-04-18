package handler

import "net/http"

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	// Read from httpOnly cookie — browser sends automatically
	refreshToken, err := getRefreshCookie(r)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "missing refresh token")
		return
	}

	// Validate against Redis
	userID, err := h.redis.Get(r.Context(), refreshToken)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid or expired refresh token")
		return
	}

	// Fetch user
	user, err := h.queries.GetUserByID(r.Context(), userID)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "user not found")
		return
	}

	// Rotate refresh token — old one deleted, new one issued
	// Prevents refresh token reuse attacks
	if err := h.redis.Delete(r.Context(), refreshToken); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to rotate session")
		return
	}

	newRefreshToken, err := generateRefreshToken()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	if err := h.redis.Save(r.Context(), newRefreshToken, userID); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to save session")
		return
	}

	// Sign new access token
	accessToken, err := h.jwt.SignAccess(userID, user.Email)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to sign token")
		return
	}

	setRefreshCookie(w, newRefreshToken)

	writeJSON(w, http.StatusOK, map[string]any{
		"access_token": accessToken,
	})
}
