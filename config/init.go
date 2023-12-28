package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

var Configuration Config

type Config struct {
	Environment string   `mapstructure:"environment"`
	Port        string   `mapstructure:"port"`
	Database    Database `mapstructure:"database"`
	BasePath    string   `mapstructure:"base_path"`
	Redis       Database `mapstructure:"redis"`
}

type Database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
}

func LoadConfig() error {
	log.Println("Starting to initialise configs")

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

	log.Println("configs initialised successfully")
	return nil
}
