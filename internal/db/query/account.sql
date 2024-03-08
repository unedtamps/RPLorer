-- name: GetAccountById :one
SELECT * FROM account WHERE id = $1;

-- name: GetAccountByEmail :one
SELECT id,email,username,password,status,role FROM account WHERE email = $1;

-- name: CreateAccount :one
INSERT INTO account ("first_name", "last_name", "username", "email", "password") 
VALUES ($1,$2,$3,$4,$5) RETURNING id, first_name, last_name, username, email,role, status;

-- name: CreateAdminAccount :one
INSERT INTO account ("first_name", "last_name", "username", "email", "password", "role", "status")
  VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id; 

-- name: CreateAccountDetail :exec
INSERT INTO account_detail ("account_id","give_like" , "get_like", "followers_count", "following_count") 
VALUES ($1,$2,$3,$4,$5);

-- name: CreateAccountFollow :exec
INSERT INTO follow ("account_followed", "account_following") VALUES ($1,$2);

-- name: UpdateGetLikeUserDetail :exec
UPDATE account_detail SET get_like = get_like + 1 WHERE account_id = $1;

-- name: UpdateDecreaseGetLikeUserDetail :exec
UPDATE account_detail SET get_like = get_like - 1 WHERE account_id = $1;

-- name: UpdateGiveLikeUserDetail :exec
UPDATE account_detail SET give_like = give_like + 1 WHERE account_id = $1;

-- name: UpdateDecreaseGiveLikeUserDetail :exec
UPDATE account_detail set give_like = give_like -1 WHERE account_id = $1;

-- name: UpdateUserStatus :exec
UPDATE account SET status = $2 WHERE id = $1;
