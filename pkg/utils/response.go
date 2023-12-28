package utils

import (
	"encoding/json"
	"net/http"
)

func ReturnJsonResponse(res http.ResponseWriter, httpCode int, msg map[string]interface{}) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)

	msgBytes, _ := json.Marshal(msg)
	res.Write(msgBytes)
}
