package handler

import "net/http"

func NewRouter(h *Handler) http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("POST /auth/register", h.Register)
    mux.HandleFunc("POST /auth/login",    h.Login)
    mux.HandleFunc("POST /auth/refresh",  h.Refresh)
    mux.HandleFunc("POST /auth/logout",   h.Logout)

    return mux
}