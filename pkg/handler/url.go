package handler

import (
	"fmt"
	"log"
	"net/http"
	"short-url/config"
	"short-url/pkg/dto/response"
	"short-url/pkg/interfaces"
	"short-url/pkg/interfaces/interactor"
	"short-url/pkg/repository"
	"short-url/pkg/usecase"
	hashFunctionInteractor "short-url/pkg/usecase/hash_function"
	"short-url/pkg/utils"
)

type URLController struct {
	URLInteractor  interactor.URLInteractor
	HashInteractor interfaces.HashInteractor
}

func NewURLController() URLController {
	return URLController{
		URLInteractor: usecase.URLInteractor{
			URL: repository.UrlRepo{},
		},
		HashInteractor: hashFunctionInteractor.RollingHash{},
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

	hashedKey, isSignedHashValue, err := uc.HashInteractor.HashFunction(r.Context(), request.LongURL)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		handlerMsg := []byte(`{"success": false, "message": Oops, something went wrong. Unable to short url}`)
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		return
	}

	request.Key = hashedKey
	request.ShortURL = fmt.Sprintf("%v/%v", config.Configuration.ShortURLBasePath, hashedKey)
	request.IsSignedKey = isSignedHashValue

	mapShortURLResponse, err := uc.URLInteractor.SaveURL(r.Context(), request)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		handlerMsg := []byte(`{"success": false, "message": Oops, something went wrong. Unable to short URL}`)
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		return
	}

	response := response.ShortURL{
		ShortURL: mapShortURLResponse.ShortURL,
		Key:      mapShortURLResponse.Key,
	}

	handlerMsg := []byte(fmt.Sprintf(`{"success": false, "message": %v}`, response))
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

	redirectUserResponse, err := uc.URLInteractor.FetchURL(r.Context(), request)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		handlerMsg := []byte(`{"success": false, "message": Oops, something went wrong. Unable to redirect user}`)
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		return
	}

	http.Redirect(w, r, redirectUserResponse.LongURL, http.StatusPermanentRedirect)
}
