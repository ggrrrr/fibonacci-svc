package system

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ggrrrr/fibonacci-svc/common/api"
)

// Help method to send response with payload and http.OK(200)
func SendPayload(w http.ResponseWriter, payload any) {
	send(w, api.Response{
		Code:    200,
		Payload: payload,
	})
}

// Help method to send response with no payload and http.InternalServerError (500)
func SendError(w http.ResponseWriter, err error) {
	send(w, api.Response{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}

// Marshals payload and writes to http writer.
func send(w http.ResponseWriter, body api.Response) {
	b, err := json.Marshal(body)
	if err != nil {
		log.Printf("unable to write response body(%v) error: %v", body, err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(body.Code)
	w.Write(b)
}
