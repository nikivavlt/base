-- name: GetTodos :many
SELECT * FROM todos
ORDER BY created_at DESC;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos (title)
VALUES ($1)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET
    title     = COALESCE($2, title),
    completed = COALESCE($3, completed)
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;
