-- name: DeleteCommentByPostId :exec
DELETE FROM comment WHERE post_id = $1;

-- name: CreateNewComment :one
INSERT INTO comment ("account_id", "post_id", "body") VALUES ($1,$2,$3) RETURNING id,body;

-- name: CreteNewCommentWithParent :one
INSERT INTO comment ("account_id", "post_id", "body", "parrent_id") VALUES ($1,$2,$3,$4) RETURNING id,body,parrent_id;

-- name: UpdateCommentCount :exec
UPDATE post SET comment_count = $2  WHERE id = $1;

-- name: UpdateCommentCountIncrement :exec
UPDATE post SET comment_count = comment_count + 1 WHERE id = $1;

-- name: UpdateCommentBody :one
UPDATE comment SET body = $1 WHERE id = $2 RETURNING id, body;

-- name: QueryCommentById :one
SELECT id, body ,account_id,status FROM comment WHERE id = $1;

-- name: QueryCommentByPost :many
SELECT id, body, parrent_id,status from comment WHERE post_id = $1 limit $2 offset $3;

-- name: QueryCommentChild :many
SELECT id , body , parrent_id ,status from comment where parrent_id = $1;

-- name: CountCommentPost :one
SELECT COUNT(*) from comment where post_id = $1;

-- name: DeleteComment :exec
UPDATE comment SET status = 'DELETED' WHERE id = $1;
