package api

import (
	"net/http"

	"github.com/ggrrrr/fibonacci-svc/common/log"
	"github.com/ggrrrr/fibonacci-svc/common/system"
)

type (
	handlers struct {
		app App
	}
)

func (a *handlers) handleNext(w http.ResponseWriter, _ *http.Request) {
	number, err := a.app.Next()
	if err != nil {
		log.Error(err).Msg("handleNext")
		system.SendError(w, err)
		return
	}
	system.SendPayload(w, number)
}

func (a *handlers) handlePrevious(w http.ResponseWriter, _ *http.Request) {
	number, err := a.app.Previous()
	if err != nil {
		log.Error(err).Msg("handlePrevious")
		system.SendError(w, err)
		return
	}
	system.SendPayload(w, number)
}

func (a *handlers) handleCurrent(w http.ResponseWriter, _ *http.Request) {
	system.SendPayload(w, a.app.Current())
}
