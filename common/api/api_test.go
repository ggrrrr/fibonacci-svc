package api

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/http"
)

func TestSend(t *testing.T) {
	testWriter := http.TestResponseWriter{}
	send(&testWriter, ApiResponse{
		Payload: "data",
		Code:    200,
		Message: "msg",
	})
	assert.Equal(t, `{"payload":"data","code":200,"message":"msg"}`, testWriter.Output)
}

func TestSendError(t *testing.T) {
	testWriter := http.TestResponseWriter{}
	SendError(&testWriter, errors.New("error msg"))
	assert.Equal(t, `{"code":500,"message":"error msg"}`, testWriter.Output)
}

func TestSendPayload(t *testing.T) {
	testWriter := http.TestResponseWriter{}
	SendPayload(&testWriter, "data")
	assert.Equal(t, `{"payload":"data","code":200}`, testWriter.Output)
}
