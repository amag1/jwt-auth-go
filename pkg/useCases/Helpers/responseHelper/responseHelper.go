package responseHelper

import (
	"auth/pkg/domain/response"
	"encoding/json"
	"net/http"
)

func ResponseBuilder(status int, message string, data interface{}) ([]byte, error) {
	response := response.Response{
		Message: message,
		Data:    data,
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return marshalResponse, nil
}

func ResponseStatusChecker(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		return
	}
}

func WriteResponse(w http.ResponseWriter, status response.Status, data interface{}) {
	response := response.Response{
		Message: status.Text,
		Data:    data,
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ResponseStatusChecker(w, []byte("500: Internal Server Error"))
		return
	}

	w.WriteHeader(status.Code)
	ResponseStatusChecker(w, marshalResponse)
}
