package service

import (
	"context"
	"errors"

	r "github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/util"
)

type UserService struct {
	*r.Store
}

type UserServiceI interface {
	CreateUser(context.Context, string, string, string) (*r.CreateUserRow, error)
	GetAllUser(context.Context, int64, int64) ([]*r.GetUsersRow, util.MetaData, error)
	LoginUser(context.Context, string, string) (*r.GetUserByEmailRow, error)
}

func newUserService(repo *r.Store) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string, email string, password string,
) (*r.CreateUserRow, error) {
	hashPassword, err := util.GenereateHasedPassword(password)
	if err != nil {
		return nil, err
	}
	user, err := s.Store.CreateUser(ctx, r.CreateUserParams{
		ID:       util.GenerateUUID(),
		Name:     name,
		Password: hashPassword,
		Email:    email,
	})
	return user, err
}

func (s *UserService) GetAllUser(
	ctx context.Context,
	page int64,
	pageSize int64,
) ([]*r.GetUsersRow, util.MetaData, error) {
	limit, offset := util.WithPagination(page, pageSize)
	getUsersParams := r.GetUsersParams{
		Limit:  limit,
		Offset: offset,
	}
	users, err := s.Queries.GetUsers(ctx, getUsersParams)
	metadata := util.WithMetadata(page, int64(len(users)), nil)
	return users, metadata, err
}

func (s *UserService) LoginUser(
	ctx context.Context,
	email string,
	password string,
) (*r.GetUserByEmailRow, error) {
	user, err := s.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	// check user not found
	if user == nil {
		return nil, errors.New("Email or Password Not Valid")
	}
	// compare password
	if ok := util.CompareHashedPassword(user.Password, password); !ok {
		return nil, errors.New("Email or Password Not Valid")
	}
	// return data
	return user, nil
}
