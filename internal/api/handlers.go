package api

import (
	"net/http"

	"github.com/ggrrrr/fibonacci-svc/common/api"
)

type (
	handlers struct {
		app App
	}
)

func (a *handlers) handleNext(w http.ResponseWriter, _ *http.Request) {
	number, err := a.app.Next()
	if err != nil {
		api.SendError(w, err)
		return
	}
	api.SendPayload(w, number)
}

func (a *handlers) handlePrevious(w http.ResponseWriter, _ *http.Request) {
	number, err := a.app.Previous()
	if err != nil {
		api.SendError(w, err)
		return
	}
	api.SendPayload(w, number)
}

func (a *handlers) handleCurrent(w http.ResponseWriter, _ *http.Request) {
	api.SendPayload(w, a.app.Current())
}
