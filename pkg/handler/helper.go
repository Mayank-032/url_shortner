package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"short-url/pkg/domain"
	"short-url/pkg/dto/request"
	"short-url/pkg/utils"
)

func validateShortURLRequest(r *http.Request) (domain.URL, error) {
	if r.Method != http.MethodPost {
		return domain.URL{}, errors.New("invalid_method")
	}

	payloadBytes, err := utils.UnmarshalRequest(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return domain.URL{}, errors.New("unable to read request body")
	}

	var urlReq request.ShortURL
	err = json.Unmarshal(payloadBytes, &urlReq)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return domain.URL{}, errors.New("unable to unmarshal request")
	}

	if len(urlReq.LongURL) == 0 {
		return domain.URL{}, errors.New("invalid_request")
	}

	return domain.URL{
		LongURL: urlReq.LongURL,
	}, nil
}

func validateRedirectURLRequest(r *http.Request) (domain.URL, error) {
	if r.Method != http.MethodGet {
		return domain.URL{}, errors.New("invalid_method")
	}

	payloadBytes, err := utils.UnmarshalRequest(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return domain.URL{}, errors.New("unable to read request body")
	}

	var urlReq request.RedirectURL
	err = json.Unmarshal(payloadBytes, &urlReq)
	if err != nil {
		log.Printf("Error: %v\n", err.Error())
		return domain.URL{}, errors.New("unable to unmarshal request")
	}

	if len(urlReq.Key) == 0 {
		return domain.URL{}, errors.New("invalid_request")
	}

	return domain.URL{
		IsKeySigned: urlReq.IsKeySigned,
		Key:         urlReq.Key,
	}, nil
}
