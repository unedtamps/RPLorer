package bootstrap

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/unedtamps/go-backend/config"
	"github.com/unedtamps/go-backend/util"
)

func connectCache() (*redis.Client, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Env.RedisHost, config.Env.RedisPort),
		Password: config.Env.RedisPassword,
		DB:       config.Env.RedisDB,
	})
	ping, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	util.Log.Infof("Redis connected: %s", ping)
	return client, nil
}
