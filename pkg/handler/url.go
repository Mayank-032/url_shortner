package handler

import (
	"database/sql"
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

func NewURLController(db *sql.DB) URLController {
	return URLController{
		URLInteractor: usecase.URLInteractor{
			URL: repository.UrlRepo{
				DB: db,
			},
		},
		HashInteractor: hashFunctionInteractor.RollingHash{},
	}
}

func (uc URLController) ShortURL(w http.ResponseWriter, r *http.Request) {
	handlerMsg := make(map[string]interface{}, 0)
	handlerMsg["status"] = false

	request, err := validateShortURLRequest(r)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		switch err.Error() {
		case "invalid_method":
			handlerMsg["message"] = "check your HTTP method: invalid http method executed"
			utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, handlerMsg)
		case "invalid_request":
			handlerMsg["message"] = "invalid Request, required params not provided. please try again"
			utils.ReturnJsonResponse(w, http.StatusBadRequest, handlerMsg)
		default:
			handlerMsg["message"] = "oops, something went wrong. please try again"
			utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		}

		return
	}

	hashedKey, isHashSigned, err := uc.HashInteractor.HashFunction(r.Context(), request.LongURL)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		handlerMsg["message"] = "oops, something went wrong. please try again"
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		return
	}

	request.Key = hashedKey
	fmt.Println("shortBasePath: " + config.Configuration.BasePath)
	request.ShortURL = fmt.Sprintf("%v/%v", config.Configuration.BasePath, hashedKey)
	request.IsKeySigned = isHashSigned

	err = uc.URLInteractor.SaveURL(r.Context(), request)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		handlerMsg["message"] = "oops, something went wrong. please try again"
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		return
	}

	response := response.ShortURL{
		ShortURL:    request.ShortURL,
		Key:         request.Key,
		IsKeySigned: isHashSigned,
	}

	handlerMsg["status"] = true
	handlerMsg["message"] = "successfully saved url mapping"
	handlerMsg["data"] = response
	utils.ReturnJsonResponse(w, http.StatusOK, handlerMsg)
	return
}

func (uc URLController) RedirectUser(w http.ResponseWriter, r *http.Request) {
	handlerMsg := make(map[string]interface{}, 0)
	handlerMsg["status"] = false

	request, err := validateRedirectURLRequest(r)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		switch err.Error() {
		case "invalid_method":
			handlerMsg["message"] = "check your HTTP method: invalid http method executed"
			utils.ReturnJsonResponse(w, http.StatusMethodNotAllowed, handlerMsg)
		case "invalid_request":
			handlerMsg["message"] = "invalid Request, required params not provided. please try again"
			utils.ReturnJsonResponse(w, http.StatusBadRequest, handlerMsg)
		default:
			handlerMsg["message"] = "oops, something went wrong. please try again"
			utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		}

		return
	}

	redirectUserResponse, err := uc.URLInteractor.FetchURL(r.Context(), request)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())

		handlerMsg["message"] = "oops, something went wrong. please try again"
		utils.ReturnJsonResponse(w, http.StatusInternalServerError, handlerMsg)
		return
	}

	http.Redirect(w, r, redirectUserResponse.LongURL, http.StatusPermanentRedirect)
}
