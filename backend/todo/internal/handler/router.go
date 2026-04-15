package handler

import (
    "net/http"
    "github.com/nikivavlt/base/todo/internal/middleware"
)

func NewRouter(h *Handler, auth *middleware.AuthMiddleware) http.Handler {
    mux := http.NewServeMux()

    mux.Handle("GET /api/todos",         auth.Protect(http.HandlerFunc(h.GetTodos)))
    mux.Handle("GET /api/todos/{id}",    auth.Protect(http.HandlerFunc(h.GetTodo)))
    mux.Handle("POST /api/todos",        auth.Protect(http.HandlerFunc(h.CreateTodo)))
    mux.Handle("PATCH /api/todos/{id}",  auth.Protect(http.HandlerFunc(h.UpdateTodo)))
    mux.Handle("DELETE /api/todos/{id}", auth.Protect(http.HandlerFunc(h.DeleteTodo)))


    return mux
}