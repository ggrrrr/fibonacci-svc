package system

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"

	"github.com/ggrrrr/fibonacci-svc/common/log"
)

type (
	Config struct {
		Addr string `envconfig:"LISTEN_ADDR"`
	}

	System struct {
		ctx         context.Context
		cancel      context.CancelFunc
		mux         *mux.Router
		addr        string
		cleanupFunc []func() error
	}
)

func NewSystem(ctx context.Context, cfg Config) *System {
	s := System{
		addr:        cfg.Addr,
		cleanupFunc: []func() error{},
	}
	s.mux = mux.NewRouter()
	s.mux.HandleFunc("/health", handlerHealth)
	s.mux.HandleFunc("/ready", handlerHealth)
	s.mux.NotFoundHandler = http.HandlerFunc(handlerNotFound)
	s.mux.MethodNotAllowedHandler = http.HandlerFunc(handlerMethodNotAllowed)
	// Add middleware for logging and tracing(otel)

	s.ctx, s.cancel = signal.NotifyContext(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return &s
}

func (s *System) AddCleanup(f func() error) {
	s.cleanupFunc = append(s.cleanupFunc, f)
}

func (s *System) Router(prefix string) *mux.Router {
	r := s.mux.PathPrefix(prefix).Subrouter()
	return r
}

// Start the listener
// this is blocking code.
func (s *System) Start() error {
	var err error
	httpd := &http.Server{
		Addr:    s.addr,
		Handler: s.mux,
	}
	log.Info().Str("addr", s.addr).Msg("Starting")
	go func() {
		err = httpd.ListenAndServe()
		if err != nil {
			log.Error(err).Str("addr", s.addr).Msg("listen")
			s.cancel()
		}
	}()
	<-s.ctx.Done()
	if err := httpd.Shutdown(s.ctx); err != nil {
		log.Error(err).Msg("httpd shutdown")
	}

	log.Info().Msg("exiting")
	s.cleanup()
	return err
}

func (s *System) cleanup() {
	for _, cleanup := range s.cleanupFunc {
		if err := cleanup(); err != nil {
			log.Error(err).Msg("cleanup call")
		}
	}
}
