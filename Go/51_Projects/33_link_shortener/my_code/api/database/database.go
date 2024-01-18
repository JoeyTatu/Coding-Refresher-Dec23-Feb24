package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	Ctx      = context.Background()
	address  = os.Getenv("DB_ADDR")
	password = os.Getenv("DB_PASS")
)

func CreateClient(databaseNumber int) *redis.Client {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       databaseNumber,
	})

	return redisDB
}
