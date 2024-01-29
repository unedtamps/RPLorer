-- name: CreateUser :one
INSERT INTO "User" (id, name, email, password) 
VALUES ($1, $2, $3, $4) 
RETURNING id, name, email;

-- name: GetUser :one
SELECT id, name, email, password 
FROM "User" WHERE id = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT id , name , email FROM "User" ORDER BY created_at ASC;

-- name: GetUsers :many
SELECT id , name , email , password 
FROM "User" ORDER BY created_at ASC LIMIT $1 OFFSET $2;

-- name: ChangeUserStatus :exec
UPDATE "User" SET account_status = $1 WHERE id = $2;

-- name: ChangeUserType :exec
UPDATE "User" SET acount_type = $1 WHERE id = $2;

-- name: DeleteUserById :exec
DELETE FROM "User" WHERE id = $1;

-- name: DeleteUserByEmail :exec
DELETE FROM "User" WHERE email = $1;
