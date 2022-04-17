package v1

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	apiResponse := ApiResponse{}

	if statusCode < http.StatusBadRequest {
		apiResponse.Status = "success"
	} else {
		apiResponse.Status = "error"
	}

	apiResponse.Result = data
	message, err := json.Marshal(apiResponse)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		message, _ = json.Marshal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(message)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(message)
	return
}
