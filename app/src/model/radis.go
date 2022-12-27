package model

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var rClient *redis.Client

func init() {
	rClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func SetKey(ctx context.Context, key string, value string) error {
	if err := rClient.Set(ctx, key, value, 6000*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func GetKey(ctx context.Context, key string) (string, error) {
	res, err := rClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}
