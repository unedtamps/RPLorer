package service

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/unedtamps/go-backend/internal/repository"
	"github.com/unedtamps/go-backend/util"

	_ "github.com/lib/pq"
)

var (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/todoapp?sslmode=disable"
)

var (
	redisHost     = "localhost"
	redisPort     = "6379"
	redisPassword = "password"
	redisDb       = 0
)

func TestMain(m *testing.M) {

	var err error
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		util.Log.Fatal("Error DB Connection: ", err)
	}
	testStore := repository.NewStore(testDB)

	redisClent := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
		DB:       redisDb,
	})
	testUserService = newUserService(testStore, redisClent)
	os.Exit(m.Run())
}
