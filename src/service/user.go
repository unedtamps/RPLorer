package service

import (
	"context"

	r "github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/util"
)

type UserService struct {
	*r.Store
}

type UserServiceI interface {
	CreateUser(context.Context, string, string, string) (*r.CreateUserRow, error)
	GetAllUser(context.Context, int64, int64) ([]*r.GetUsersRow, util.MetaData, error)
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
