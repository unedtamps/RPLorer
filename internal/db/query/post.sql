-- name: CreatePost :one
INSERT INTO post ("account_id", "caption", "likes_count", "comment_count")
VALUES ($1,$2,$3,$4) RETURNING id;

-- name: CreatePostImages :exec
INSERT INTO images ("post_id", "img_url") VALUES ($1,$2);

-- name: CreateLikedPost :exec
INSERT INTO liked ("account_id", "post_id") VALUES ($1,$2);

-- name: DeleteLikedPost :exec
DELETE FROM liked WHERE account_id = $1 AND post_id = $2;

-- name: CreateComment :one
INSERT INTO comment ("account_id", "post_id", "body", "parrent_id") VALUES ($1,$2,$3,$4) RETURNING id;


-- name: UpdateLikeCountDecrement :one
UPDATE post SET likes_count = likes_count - 1 WHERE id = $1 AND likes_count > 0 AND status = 'OK' RETURNING account_id;

-- name: UpdateLikeCountIncrement :one
UPDATE post SET likes_count = likes_count + 1 WHERE id = $1 AND status = 'OK' RETURNING account_id;

-- name: UpdatePostDetailCount :exec
UPDATE post SET likes_count = $1, comment_count = $2 WHERE id = $3;

-- name: UpdateFollowersCount :exec
UPDATE account_detail SET followers_count = followers_count + 1 WHERE account_id = $1;

-- name: UpdateFollowingCount :exec
UPDATE account_detail SET following_count = following_count + 1 WHERE account_id = $1;

-- name: InsertNewPost :one
INSERT INTO post ("account_id", "caption") VALUES ($1,$2) RETURNING id, caption;

-- name: InsertNewPostImages :exec
INSERT INTO images ("id", "post_id", "img_url") VALUES ($1,$2,$3);

-- name: QueryPostByUserID :many
SELECT id, caption FROM post WHERE account_id = $1 AND status = 'OK' LIMIT $2 OFFSET $3 ;

-- name: QueryImageByPostID :many
SELECT i.img_url FROM images i JOIN post p on p.id = i.post_id AND p.status = 'OK' WHERE i.post_id = $1;

-- name: CoutAccoutPost :one
SELECT COUNT(*) FROM post WHERE account_id = $1 AND status = 'OK';

-- name: QueryPostandImages :many
SELECT p.id as post_id, p.caption,p.likes_count, p.comment_count, i.img_url FROM post p 
LEFT JOIN images i ON p.id = i.post_id 
WHERE p.id = $1 AND p.status = 'OK';

-- name: ChangePostStatus :exec
UPDATE post SET status = $1 WHERE id = $2 AND account_id = $3;

-- name: UpdatePostCaption :one
UPDATE post SET caption = $1 WHERE id = $2 RETURNING id, caption, account_id;

-- name: QueryUserIdfromPost :one
SELECT account_id FROM post WHERE id = $1;

-- name: DeletePostById :exec
DELETE FROM post WHERE id = $1;

-- name: QueryGetAccoutFromLikedByPostId :many
SELECT account_id FROM liked WHERE post_id = $1; 

-- name: DeleteImageByPostId :exec
DELETE FROM images WHERE post_id = $1;

-- name: DeleteLikedByPostId :exec
DELETE FROM liked WHERE post_id = $1;

