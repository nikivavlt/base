package handler

import (
    "net/http"
)

func NewRouter(h *Handler) http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /api/todos",        h.GetTodos)
    mux.HandleFunc("GET /api/todos/{id}",   h.GetTodo)
    mux.HandleFunc("POST /api/todos",       h.CreateTodo)
    mux.HandleFunc("PATCH /api/todos/{id}", h.UpdateTodo)
    mux.HandleFunc("DELETE /api/todos/{id}",h.DeleteTodo)

    return mux
}