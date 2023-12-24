package handler

import (
	"fmt"
	"log"
	"net/http"
	"short-url/pkg/interfaces/interactor"
	"short-url/pkg/repository"
	"short-url/pkg/usecase"
	"short-url/pkg/utils"
)

type URLController struct {
	URLInteractor interactor.URLInteractor
}

func NewURLController() URLController {
	return URLController{
		URLInteractor: usecase.URLInteractor{
			URL: repository.UrlRepo{},
		},
	}
}

func (uc URLController) ShortURL(w http.ResponseWriter, r *http.Request) {
	request, err := validateShortURLRequest(r)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		switch err.Error() {
		case "invalid_method":
			handlerMsg := []byte(`{"success": false, "message": "Check your HTTP method: Invalid HTTP method executed"}`)
			utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, handlerMsg)
		case "invalid_request":
			handlerMsg := []byte(`{"success": false, "message": "Invalid Request, required params not provided. Please try again"}`)
			utils.ReturnJsonResponse(w, http.StatusBadRequest, handlerMsg)
		default:
			handlerMsg := []byte(`{"success": false, "message": "Oops, something went wrong. Please try again"}`)
			utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		}

		return
	}

	mapShortURLResponse, err := uc.URLInteractor.ShortURLMapper(r.Context(), request)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		handlerMsg := []byte(`{"success": false, "message": Oops, something went wrong. Unable to redirect user}`)
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		return
	}

	handlerMsg := []byte(fmt.Sprintf(`{"success": false, "message": %v}`, mapShortURLResponse))
	utils.ReturnJsonResponse(w, http.StatusOK, handlerMsg)
	return
}

func (uc URLController) RedirectUser(w http.ResponseWriter, r *http.Request) {
	request, err := validateRedirectURLRequest(r)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		switch err.Error() {
		case "invalid_method":
			handlerMsg := []byte(`{"success": false, "message": "Check your HTTP method: Invalid HTTP method executed"}`)
			utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, handlerMsg)
		case "invalid_request":
			handlerMsg := []byte(`{"success": false, "message": "Invalid Request, required params not provided. Please try again"}`)
			utils.ReturnJsonResponse(w, http.StatusBadRequest, handlerMsg)
		default:
			handlerMsg := []byte(`{"success": false, "message": "Oops, something went wrong. Please try again"}`)
			utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		}

		return
	}

	redirectUserResponse, err := uc.URLInteractor.FetchLongURL(r.Context(), request)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		handlerMsg := []byte(`{"success": false, "message": Oops, something went wrong. Unable to redirect user}`)
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		return
	}

	http.Redirect(w, r, redirectUserResponse.LongURL, http.StatusPermanentRedirect)
}
