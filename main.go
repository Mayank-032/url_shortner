package main

import (
	"fmt"
	"log"
	"net/http"
	"short-url/infrastructure/config"
	"short-url/infrastructure/database"
	"short-url/infrastructure/routes"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: %v\n. unable to load config", err.Error())
		return
	}

	err = database.InitMySQL()
	if err != nil {
		log.Fatalf("Error: %v\n. unable to connect to db", err.Error())
		return
	}

	err = database.InitRedis()
	if err != nil {
		log.Fatalf("Error: %v\n. unable to connect to db", err.Error())
		return
	}

	r := http.NewServeMux()
	routes.InitRoutes(r)

	port := fmt.Sprintf(":%v", config.Configuration.Port)
	fmt.Println("Starting to listen on PORT " + port)
	if err = http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Error: %v\n. server shutdown gracefully", err.Error())
		return
	}
}
