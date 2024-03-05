package service

import (
	"github.com/redis/go-redis/v9"
	r "github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/src/helper"
	"gopkg.in/gomail.v2"
)

type Service struct {
	Account AccountService
}

func NewService(repo *r.Store, cache *redis.Client, d *gomail.Dialer) *Service {
	helper.Dialer = d
	return &Service{
		Account: *newAccountService(repo, cache),
	}
}
