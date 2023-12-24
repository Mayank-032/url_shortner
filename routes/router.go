package routes

import (
	"net/http"
	"short-url/pkg/handler"
)

func InitRoutes(router *http.ServeMux) {
	urlController := handler.NewURLController()
	router.HandleFunc("url/short", urlController.ShortURL)
	router.HandleFunc("url/redirect", urlController.RedirectUser)
}
