package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	r "github.com/unedtamps/go-backend/internal/repository"
)

type AccountService struct {
	repo  *r.Store
	cache *redis.Client
}

type AccountServiceI interface {
	GetOne(c *gin.Context) (*r.Account, error)
}

func newAccountService(repo *r.Store, cache *redis.Client) *AccountService {
	return &AccountService{repo, cache}
}

func (s *AccountService) GetOne(c *gin.Context) (*r.Account, error) {
	acc, err := s.repo.SelectAccount(c, uuid.MustParse("d06d41d0-c01f-4970-a1b1-b1813b641fec"))
	if err != nil {
		return nil, err
	}
	return acc, err
}
