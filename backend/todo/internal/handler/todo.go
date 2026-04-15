package handler

import (
    "database/sql"
    "encoding/json"
    "errors"
    "net/http"
    "strconv"

    "github.com/nikivavlt/base/todo/internal/db"
)

type Handler struct {
    queries *db.Queries
}

func New(queries *db.Queries) *Handler {
    return &Handler{queries: queries}
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
    todos, err := h.queries.GetTodos(r.Context())
    if err != nil {
        writeError(w, http.StatusInternalServerError, "failed to fetch todos")
        return
    }

    if todos == nil {
        todos = []db.Todo{}
    }

    writeJSON(w, http.StatusOK, todos)
}

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
    if err != nil {
        writeError(w, http.StatusBadRequest, "invalid id")
        return
    }

    todo, err := h.queries.GetTodo(r.Context(), id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            writeError(w, http.StatusNotFound, "todo not found")
            return
        }
        writeError(w, http.StatusInternalServerError, "failed to fetch todo")
        return
    }

    writeJSON(w, http.StatusOK, todo)
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
    var body struct {
        Title string `json:"title"`
    }

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        writeError(w, http.StatusBadRequest, "invalid request body")
        return
    }

    if body.Title == "" {
        writeError(w, http.StatusBadRequest, "title is required")
        return
    }

    todo, err := h.queries.CreateTodo(r.Context(), body.Title)
    if err != nil {
        writeError(w, http.StatusInternalServerError, "failed to create todo")
        return
    }

    writeJSON(w, http.StatusCreated, todo)
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
    if err != nil {
        writeError(w, http.StatusBadRequest, "invalid id")
        return
    }

    var body struct {
        Title     *string `json:"title"`
        Completed *bool   `json:"completed"`
    }

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        writeError(w, http.StatusBadRequest, "invalid request body")
        return
    }

    todo, err := h.queries.UpdateTodo(r.Context(), db.UpdateTodoParams{
        ID:        id,
        Title:     body.Title,
        Completed: body.Completed,
    })
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            writeError(w, http.StatusNotFound, "todo not found")
            return
        }
        writeError(w, http.StatusInternalServerError, "failed to update todo")
        return
    }

    writeJSON(w, http.StatusOK, todo)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
    if err != nil {
        writeError(w, http.StatusBadRequest, "invalid id")
        return
    }

    if err := h.queries.DeleteTodo(r.Context(), id); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            writeError(w, http.StatusNotFound, "todo not found")
            return
        }
        writeError(w, http.StatusInternalServerError, "failed to delete todo")
        return
    }

    w.WriteHeader(http.StatusNoContent)
}