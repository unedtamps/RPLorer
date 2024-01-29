package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/unedtamps/go-backend/util"
)

func createRandomUser(t *testing.T) CreateUserRow {
  args := CreateUserParams{
    ID: util.RandomUUID(),
    Name: util.RandomString(),
    Email: util.RandomEmail(),
    Password: util.RandomHashedPassword(),
  }
  user , err :=  testQueries.CreateUser(context.Background(),args)
  require.NoError(t,err)
  require.NotEmpty(t,user)
  require.Equal(t,args.ID, user.ID)
  require.Equal(t,args.Name, user.Name)
  require.Equal(t,args.Email, user.Email)

  return *user
}

func TestCreateUser(t *testing.T) {
  createRandomUser(t)
}

func TestGetUser(t *testing.T) {
  user1 := createRandomUser(t)
  user2, err := testQueries.GetUser(context.Background(),user1.ID)
  require.NoError(t,err)
  require.NotEmpty(t,user2)
  require.Equal(t,user1.ID, user2.ID)
  require.Equal(t,user1.Name, user2.Name)
  require.Equal(t,user1.Email, user2.Email)
}

func TestDeleteUserById(t *testing.T) {
  user1 := createRandomUser(t)
  err := testQueries.DeleteUserById(context.Background(),user1.ID)
  require.NoError(t,err)
  user2 , err := testQueries.GetUser(context.Background(),user1.ID)
  require.Error(t,err)
  require.EqualError(t,err,sql.ErrNoRows.Error())
  require.Empty(t,user2)
}

func TestDeleteUserByEmail(t *testing.T) {
  user1 := createRandomUser(t)
  err := testQueries.DeleteUserByEmail(context.Background(),user1.Email)
  require.NoError(t,err)
  user2, err := testQueries.GetUser(context.Background(),user1.ID)
  require.Error(t,err)
  require.EqualError(t,err,sql.ErrNoRows.Error())
  require.Empty(t,user2)
}

func TestChangeUserStatus(t *testing.T) {
  user1 := createRandomUser(t)
  paramUserStatus := ChangeUserStatusParams{
    AccountStatus: sql.NullBool{Bool:true, Valid: true},
    ID: user1.ID,
  }

  err := testQueries.ChangeUserStatus(context.Background(),paramUserStatus)
  require.NoError(t,err)
}
