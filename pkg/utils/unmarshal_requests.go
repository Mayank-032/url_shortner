package utils

import (
	"io"
	"log"
)

func UnmarshalRequest(body io.ReadCloser) ([]byte, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		log.Println("Error: " + err.Error())
		return nil, err
	}

	return bodyBytes, nil
}
