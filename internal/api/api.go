package api

//go:generate mockgen -source=$GOFILE -destination=app_mock.go -package=api

import (
	"github.com/ggrrrr/fibonacci-svc/internal/fi"
	"github.com/gorilla/mux"
)

type (
	App interface {
		Next() (fi.Number, error)
		Previous() (fi.Number, error)
		Current() fi.Number
	}
)

func Register(rootMux *mux.Router, app App) {
	h := handlers{
		app: app,
	}

	rootMux.HandleFunc("/next", h.handleNext)

	rootMux.HandleFunc("/previous", h.handlePrevious).Methods("GET")
	rootMux.HandleFunc("/prev", h.handlePrevious).Methods("GET")

	rootMux.HandleFunc("/current", h.handleCurrent).Methods("GET")
	rootMux.HandleFunc("/cur", h.handleCurrent).Methods("GET")
}
