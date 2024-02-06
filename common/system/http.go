package system

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ggrrrr/fibonacci-svc/common/log"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	send(w, ApiResponse{
		Code:    http.StatusOK,
		Message: "ok",
	})
}

func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	log.Error(errors.New("StatusNotFound")).Any("asd", r.Method).Any("path", r.URL.Path).Send()
	send(w, ApiResponse{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("path not found: [%s] %s", r.Method, r.URL.Path),
	})
}

func handlerMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	log.Error(errors.New("StatusMethodNotAllowed")).Any("asd", r.Method).Any("path", r.URL.Path).Send()
	send(w, ApiResponse{
		Code:    http.StatusMethodNotAllowed,
		Message: fmt.Sprintf("method not allowed: [%s] %s", r.Method, r.URL.Path),
	})
}
