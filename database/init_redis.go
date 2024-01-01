package database

import (
	"fmt"
	"log"
	"short-url/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	// connectionString := fmt.Sprintf("redis://%v:%v@%v:%v", config.Configuration.Redis.Username, config.Configuration.Redis.Password, config.Configuration.Redis.Host, config.Configuration.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Configuration.Redis.Host, config.Configuration.Redis.Port),
		Username: config.Configuration.Redis.Username,
		Password: config.Configuration.Redis.Password,
		DB:       0,
	})

	RedisClient = client
	log.Println("successfuly connected with redis")
	return nil
}
