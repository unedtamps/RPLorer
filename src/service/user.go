package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	r "github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/util"
)

type UserService struct {
	*r.Store
	cache *redis.Client
}

type UserServiceI interface {
	CreateUser(context.Context, string, string, string) (*r.CreateUserRow, error)
	GetAllUser(context.Context, int64, int64) ([]*r.GetUsersRow, util.MetaData, error)
	LoginUser(context.Context, string, string) (*r.GetUserByEmailRow, error)
	GetUserByEmail(context.Context, string) (*r.GetUserByEmailRow, error)
}

func newUserService(repo *r.Store, cache *redis.Client) *UserService {
	return &UserService{repo, cache}
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

	if err != nil {
		switch err.Error() {
		case "pq: duplicate key value violates unique constraint \"User_email_key\"":
			return nil, errors.New("Account Already Exist")
		default:
			return nil, err
		}
	}
	return user, nil
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

func (s *UserService) GetUserByEmail(
	ctx context.Context,
	email string,
) (*r.GetUserByEmailRow, error) {

	val, err := s.cache.Get(ctx, email).Result()

	if err == nil {
		res, err := util.ParseCache[r.GetUserByEmailRow](val)
		return res, err
	}

	user, err := s.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	redisData, _ := json.Marshal(user)
	err = s.cache.Set(ctx, email, redisData, 0).Err()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) LoginUser(
	ctx context.Context,
	email string,
	password string,
) (*r.GetUserByEmailRow, error) {

	val, err := s.cache.Get(ctx, email).Result()

	if err == nil {
		res, err := util.ParseCache[r.GetUserByEmailRow](val)
		if ok := util.CompareHashedPassword(res.Password, password); !ok {
			return nil, errors.New("Email or Password Not Valid")
		}
		return res, err
	}

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

	redisData, _ := json.Marshal(user)
	err = s.cache.Set(ctx, email, redisData, time.Hour).Err()
	// return data
	return user, err
}
