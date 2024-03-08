package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	r "github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/util"
)

type AccountService struct {
	repo  *r.Store
	cache *redis.Client
}

type AccountServiceI interface {
	CreateUserService(
		ctx context.Context,
		firstName, lastName, username, email, password string,
	) (*r.CreateAccountRow, *ErrorService)
	ActivatedAccount(ctx context.Context, id string) *ErrorService
	ReSendEmailConfirmation(ctx context.Context, email string) *ErrorService
	LoginService(ctx context.Context, email, password string) (*string, *ErrorService)
}

func newAccountService(repo *r.Store, cache *redis.Client) AccountServiceI {
	return &AccountService{repo, cache}
}

func newCreateUserParam(
	firstName, lastName, username, email, password string,
) r.CreateAccountParams {
	return r.CreateAccountParams{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
	}
}

func (a *AccountService) CreateUserService(
	ctx context.Context,
	firstName, lastName, username, email, password string,
) (*r.CreateAccountRow, *ErrorService) {
	// hash password
	hash, err := util.GenereateHasedPassword(password)
	if err != nil {
		return nil, newError(err, 500)
	}

	account, err := a.repo.CreateAccount(
		ctx,
		newCreateUserParam(firstName, lastName, username, email, hash),
	)
	if err != nil {
		return nil, newError(err, 500)
	}
	err = a.repo.CreateAccountDetail(ctx, r.CreateAccountDetailParams{
		AccountID:      account.ID,
		GiveLike:       0,
		GetLike:        0,
		FollowersCount: 0,
		FollowingCount: 0,
	})
	if err != nil {
		return nil, newError(err, 500)
	}

	// createJwtToken
	token, err := makeJwtToken(
		account.ID.String(),
		account.Email,
		account.Username,
		string(account.Role),
	)
	if err != nil {
		return nil, newError(err, 500)
	}
	// send email confirmation
	go sendEmailConfirmation(account.ID.String(), token, account.Email, account.Username)

	if err != nil {
		return nil, newError(err, 500)
	}
	return account, nil
}

func (a *AccountService) ActivatedAccount(ctx context.Context, id string) *ErrorService {
	err := a.repo.UpdateUserStatus(ctx, r.UpdateUserStatusParams{
		ID:     uuid.MustParse(id),
		Status: r.AccountStatusACTIVE,
	})
	if err != nil {
		return newError(err, 500)
	}
	return nil
}

func (a *AccountService) ReSendEmailConfirmation(ctx context.Context, email string) *ErrorService {
	acc, err := a.repo.GetAccountByEmail(ctx, email)
	if err != nil {
		return newError(err, 500)
	}
	if acc.Status == r.AccountStatusACTIVE {
		return newError(errors.New("account already activated"), 403)
	}
	token, err := makeJwtToken(acc.ID.String(), acc.Email, acc.Username, string(acc.Role))
	if err != nil {
		return newError(err, 500)
	}
	go sendEmailConfirmation(acc.ID.String(), token, acc.Email, acc.Username)
	return nil
}

func (a *AccountService) LoginService(
	ctx context.Context,
	email, password string,
) (*string, *ErrorService) {
	acc, err := a.repo.GetAccountByEmail(ctx, email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, newError(errors.New("email or password not found"), 404)
		default:
			return nil, newError(err, 500)
		}
	}
	if acc.Status != r.AccountStatusACTIVE {
		return nil, newError(errors.New("account not activated"), 403)
	}

	if ok := util.CompareHashedPassword(acc.Password, password); !ok {
		return nil, newError(errors.New("email or password not found"), 404)
	}
	token, err := makeJwtToken(acc.ID.String(), acc.Email, acc.Username, string(acc.Role))
	if err != nil {
		return nil, newError(err, 500)
	}

	return &token, nil
}
