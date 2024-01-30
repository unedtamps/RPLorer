-- name: CreateTodo :one 
INSERT INTO "Todo" (id, user_id,todo_name, description)
VALUES ($1, $2, $3, $4) RETURNING id,todo_name, description, status;

-- name: GetTodoByUserId :many
SELECT todo_name, status, description 
FROM "Todo" 
WHERE user_id = $1
LIMIT $2 OFFSET $3;
