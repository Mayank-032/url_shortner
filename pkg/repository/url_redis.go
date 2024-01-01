package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func (ur UrlRepo) Get(ctx context.Context, key string) (string, error) {
	var value string // Change the type from *string to string

	value, err := ur.RedisClient.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		if err == redis.Nil {
			return "", nil
		}
		return "", errors.New("redis error while fetching info")
	}

	if value == "" {
		return "", nil // or handle the empty value case differently
	}
	return value, nil
}

func (ur UrlRepo) Set(ctx context.Context, key, value string) {
	ur.RedisClient.Set(ctx, key, value, time.Hour*24)
}
