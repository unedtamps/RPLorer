-- name: SelectAccount :one
SELECT * FROM account WHERE id = $1;
