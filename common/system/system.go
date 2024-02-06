package system

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ggrrrr/fibonacci-svc/common/log"
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
	s.mux.HandleFunc("/health", handlerHealth)
	s.mux.HandleFunc("/ready", handlerHealth)
	s.mux.NotFoundHandler = http.HandlerFunc(handlerNotFound)
	s.mux.MethodNotAllowedHandler = http.HandlerFunc(handlerMethodNotAllowed)
	return &s
}

func (s *System) Router() *mux.Router {
	r := s.mux.PathPrefix("/v1").Subrouter()
	return r
}

// Start the listener
// this is blocking code.
func (s *System) Start() error {
	log.Info().Str("addr", s.addr).Msg("mount")
	return http.ListenAndServe(s.addr, s.mux)
}
