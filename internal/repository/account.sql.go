// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: account.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO account ("first_name", "last_name", "username", "email", "password") 
VALUES ($1,$2,$3,$4,$5) RETURNING id, first_name, last_name, username, email,role, status
`

type CreateAccountParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateAccountRow struct {
	ID        uuid.UUID     `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Role      Role          `json:"role"`
	Status    AccountStatus `json:"status"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (*CreateAccountRow, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.Email,
		arg.Password,
	)
	var i CreateAccountRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Role,
		&i.Status,
	)
	return &i, err
}

const createAccountDetail = `-- name: CreateAccountDetail :exec
INSERT INTO account_detail ("account_id","give_like" , "get_like", "followers_count", "following_count") 
VALUES ($1,$2,$3,$4,$5)
`

type CreateAccountDetailParams struct {
	AccountID      uuid.UUID `json:"account_id"`
	GiveLike       int64     `json:"give_like"`
	GetLike        int64     `json:"get_like"`
	FollowersCount int64     `json:"followers_count"`
	FollowingCount int64     `json:"following_count"`
}

func (q *Queries) CreateAccountDetail(ctx context.Context, arg CreateAccountDetailParams) error {
	_, err := q.db.ExecContext(ctx, createAccountDetail,
		arg.AccountID,
		arg.GiveLike,
		arg.GetLike,
		arg.FollowersCount,
		arg.FollowingCount,
	)
	return err
}

const createAccountFollow = `-- name: CreateAccountFollow :exec
INSERT INTO follow ("account_followed", "account_following") VALUES ($1,$2)
`

type CreateAccountFollowParams struct {
	AccountFollowed  uuid.UUID `json:"account_followed"`
	AccountFollowing uuid.UUID `json:"account_following"`
}

func (q *Queries) CreateAccountFollow(ctx context.Context, arg CreateAccountFollowParams) error {
	_, err := q.db.ExecContext(ctx, createAccountFollow, arg.AccountFollowed, arg.AccountFollowing)
	return err
}

const createAdminAccount = `-- name: CreateAdminAccount :one
INSERT INTO account ("first_name", "last_name", "username", "email", "password", "role", "status")
  VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id
`

type CreateAdminAccountParams struct {
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Role      Role          `json:"role"`
	Status    AccountStatus `json:"status"`
}

func (q *Queries) CreateAdminAccount(ctx context.Context, arg CreateAdminAccountParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createAdminAccount,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Role,
		arg.Status,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT id,email,username,password,status,role FROM account WHERE email = $1
`

type GetAccountByEmailRow struct {
	ID       uuid.UUID     `json:"id"`
	Email    string        `json:"email"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Status   AccountStatus `json:"status"`
	Role     Role          `json:"role"`
}

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (*GetAccountByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getAccountByEmail, email)
	var i GetAccountByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Username,
		&i.Password,
		&i.Status,
		&i.Role,
	)
	return &i, err
}

const getAccountById = `-- name: GetAccountById :one
SELECT id, first_name, last_name, username, email, role, status, password, created_at, updated_at FROM account WHERE id = $1
`

func (q *Queries) GetAccountById(ctx context.Context, id uuid.UUID) (*Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountById, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Role,
		&i.Status,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updateDecreaseGetLikeUserDetail = `-- name: UpdateDecreaseGetLikeUserDetail :exec
UPDATE account_detail SET get_like = get_like - 1 WHERE account_id = $1
`

func (q *Queries) UpdateDecreaseGetLikeUserDetail(ctx context.Context, accountID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, updateDecreaseGetLikeUserDetail, accountID)
	return err
}

const updateDecreaseGiveLikeUserDetail = `-- name: UpdateDecreaseGiveLikeUserDetail :exec
UPDATE account_detail set give_like = give_like -1 WHERE account_id = $1
`

func (q *Queries) UpdateDecreaseGiveLikeUserDetail(ctx context.Context, accountID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, updateDecreaseGiveLikeUserDetail, accountID)
	return err
}

const updateGetLikeUserDetail = `-- name: UpdateGetLikeUserDetail :exec
UPDATE account_detail SET get_like = get_like + 1 WHERE account_id = $1
`

func (q *Queries) UpdateGetLikeUserDetail(ctx context.Context, accountID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, updateGetLikeUserDetail, accountID)
	return err
}

const updateGiveLikeUserDetail = `-- name: UpdateGiveLikeUserDetail :exec
UPDATE account_detail SET give_like = give_like + 1 WHERE account_id = $1
`

func (q *Queries) UpdateGiveLikeUserDetail(ctx context.Context, accountID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, updateGiveLikeUserDetail, accountID)
	return err
}

const updateUserStatus = `-- name: UpdateUserStatus :exec
UPDATE account SET status = $2 WHERE id = $1
`

type UpdateUserStatusParams struct {
	ID     uuid.UUID     `json:"id"`
	Status AccountStatus `json:"status"`
}

func (q *Queries) UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateUserStatus, arg.ID, arg.Status)
	return err
}