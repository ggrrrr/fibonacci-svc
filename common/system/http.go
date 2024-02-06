package system

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ggrrrr/fibonacci-svc/common/log"
	"github.com/gorilla/mux"
)

type (
	System struct {
		mux  *mux.Router
		addr string
	}
)

func NewSystem() *System {
	s := System{
		addr: "localhost:8090",
	}
	s.mux = mux.NewRouter()
	s.mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		send(w, ApiResponse{Code: 200, Payload: "asd"})
	})
	s.mux.HandleFunc("/health", handlerHealth)
	s.mux.HandleFunc("/ready", handlerHealth)
	s.mux.NotFoundHandler = http.HandlerFunc(handlerNotFound)
	s.mux.MethodNotAllowedHandler = http.HandlerFunc(handlerMethodNotAllowed)
	// s.mux.MethodNotAllowed(MethodNotAllowedHandler)
	return &s
}

func (s *System) Router() *mux.Router {
	r := s.mux.PathPrefix("/v1").Subrouter()
	// s.mux.Use(r)
	return r
}

func (s *System) Mount(prefix string, mux http.Handler) {
	log.Info().Str("prefix", prefix).Msg("mount")
	// r := s.mux.PathPrefix(prefix).Subrouter()
	s.mux.Handle("/", mux)
}

func (s *System) Start() error {
	log.Info().Str("addr", s.addr).Msg("mount")
	return http.ListenAndServe(s.addr, s.mux)
}

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
