package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

var Configuration Config

type Config struct {
	Environment string        `json:"environment"`
	Port        string        `json:"port"`
	Database    MySQLDatabase `json:"database"`
}

type MySQLDatabase struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Schema   string `json:"schema"`
}

func LoadConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
    if err != nil {
		log.Println("Error: " + err.Error())
        return errors.New("unable to read config file")
    }

	err = viper.Unmarshal(&Configuration)
	if err != nil {
		log.Println("Error: " + err.Error())
		return errors.New("unable to unmarshal file data to struct")
	}

	return nil
}