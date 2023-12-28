package routes

import (
	"log"
	"net/http"
	"short-url/database"
	"short-url/pkg/handler"
)

func InitRoutes(router *http.ServeMux) {
	log.Println("Starting to initialise routes")

	urlController := handler.NewURLController(database.DB)

	router.HandleFunc("/url/short", urlController.ShortURL)
	router.HandleFunc("/url/redirect", urlController.RedirectUser)

	log.Println("Routes initialised successfully")
}
