package database

import (
	"fmt"
	"log"
	"short-url/config"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() error {
	connectionString := fmt.Sprintf("redis://%v:%v@%v:%v/%v", config.Configuration.Redis.Username, config.Configuration.Redis.Password, config.Configuration.Redis.Host, config.Configuration.Redis.Port, config.Configuration.Redis.Schema)

	opt, err := redis.ParseURL(connectionString)
	if err != nil {
		log.Println("Error: " + err.Error())
		return err
	}
	client := redis.NewClient(opt)

	RedisClient = client
	log.Println("successfuly connected with redis")
	return nil
}
