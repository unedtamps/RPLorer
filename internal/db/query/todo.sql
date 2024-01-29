-- name: CreateTodo :one 
INSERT INTO "Todo" (id, user_id,todo_name, description)
VALUES ($1, $2, $3, $4) RETURNING id, user_id,todo_name, description, status;
