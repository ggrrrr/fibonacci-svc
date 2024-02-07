package system

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ggrrrr/fibonacci-svc/common/log"
)

type (
	Config struct {
		Addr string `envconfig:"LISTEN_ADDR"`
	}

	System struct {
		mux  *mux.Router
		addr string
	}
)

func NewSystem(cfg Config) *System {
	s := System{
		addr: cfg.Addr,
	}
	s.mux = mux.NewRouter()
	s.mux.HandleFunc("/health", handlerHealth)
	s.mux.HandleFunc("/ready", handlerHealth)
	s.mux.NotFoundHandler = http.HandlerFunc(handlerNotFound)
	s.mux.MethodNotAllowedHandler = http.HandlerFunc(handlerMethodNotAllowed)
	// Add middleware for logging and tracing(otel)
	return &s
}

func (s *System) Router(prefix string) *mux.Router {
	r := s.mux.PathPrefix(prefix).Subrouter()
	return r
}

// Start the listener
// this is blocking code.
func (s *System) Start() error {
	log.Info().Str("addr", s.addr).Msg("Starting")
	err := http.ListenAndServe(s.addr, s.mux)
	if err != nil {
		log.Error(err).Str("addr", s.addr).Msg("listen")
		return err
	}
	return nil
}
