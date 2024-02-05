package api

//go:generate mockgen -source=$GOFILE -destination=app_mock.go -package=api

import (
	"net/http"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
)

type (
	App interface {
		Next() (fi.Number, error)
		Previous() (fi.Number, error)
		Current() fi.Number
	}
)

func Register(rootMux *http.ServeMux, app App) {
	h := handlers{
		app: app,
	}
	rootMux.HandleFunc("/next", h.handleNext)
	rootMux.HandleFunc("/previous", h.handlePrevious)
	rootMux.HandleFunc("/current", h.handleCurrent)
}
