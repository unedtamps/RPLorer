package service

import (
	"github.com/redis/go-redis/v9"
	"github.com/unedtamps/go-backend/config"
	r "github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/src/helper"
	"github.com/unedtamps/go-backend/src/middleware"
	"github.com/unedtamps/go-backend/util"
	"gopkg.in/gomail.v2"
)

type Service struct {
	Account AccountServiceI
	Post    PostServiceI
}
type ErrorService struct {
	Error error
	Code  int
}

func newError(err error, code int) *ErrorService {
	return &ErrorService{err, code}
}

func NewService(repo *r.Store, cache *redis.Client, d *gomail.Dialer) *Service {
	helper.Dialer = d
	return &Service{
		Account: newAccountService(repo, cache),
		Post:    newPostService(repo, cache),
	}
}

func makeJwtToken(id, email, name, role string) (string, error) {
	return middleware.CreateJwtToken(middleware.Credentials{
		Id:    id,
		Email: email,
		Name:  name,
		Role:  role,
	})
}

func sendEmailConfirmation(
	id, token, email, name string,
) {
	host := config.Env.ServerHost
	//send email confirmation
	body := make(chan string)
	go func() {
		body <- util.ParseAccountConfirmation(util.EmailConfirm{
			Name:  name,
			Token: token,
			Host:  host,
		})
	}()
	go helper.NewEmail("Account Confirmation", email, body).Send()
}
