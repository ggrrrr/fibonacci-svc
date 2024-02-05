package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	ApiResponse struct {
		Payload any    `json:"payload,omitempty"`
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}
)

func SendPayload(w http.ResponseWriter, payload any) {
	send(w, ApiResponse{
		Code:    200,
		Payload: payload,
	})
}

func SendError(w http.ResponseWriter, err error) {
	send(w, ApiResponse{
		Code:    500,
		Message: err.Error(),
	})
}

func send(w http.ResponseWriter, body ApiResponse) {
	b, err := json.Marshal(body)
	if err != nil {
		log.Printf("unable to write response body(%v) error: %v", body, err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(body.Code)
	w.Write(b)
}
