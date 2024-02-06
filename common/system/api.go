package system

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	// We use this model as generic response for all HTTP calls to this service.
	// including errors
	ApiResponse struct {
		Payload any    `json:"payload,omitempty"`
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}
)

// Help method to send response with payload and http.OK(200)
func SendPayload(w http.ResponseWriter, payload any) {
	send(w, ApiResponse{
		Code:    200,
		Payload: payload,
	})
}

// Help method to send response with no payload and http.InternalServerError (500)
func SendError(w http.ResponseWriter, err error) {
	send(w, ApiResponse{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}

// Marshals payload and writes to http writer.
func send(w http.ResponseWriter, body ApiResponse) {
	b, err := json.Marshal(body)
	if err != nil {
		log.Printf("unable to write response body(%v) error: %v", body, err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(body.Code)
	w.Write(b)
}
