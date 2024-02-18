package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/unedtamps/go-backend/util"
)

var testUserService *UserService

func TestLogin(t *testing.T) {
	ctx := context.Background()
	name := util.RandomString()
	email := util.RandomEmail()
	password := util.RandomString()
	user, err := testUserService.CreateUser(ctx, name, email, password)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, name, user.Name)
	require.Equal(t, email, user.Email)
	require.NotEmpty(t, user.ID)

	userLogin, err := testUserService.LoginUser(ctx, email, password)
	require.NoError(t, err)
	require.NotEmpty(t, userLogin)
	require.Equal(t, user.Email, userLogin.Email)
	require.Equal(t, user.ID, userLogin.ID)
}
