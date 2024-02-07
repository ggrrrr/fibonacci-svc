package system

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/http"

	"github.com/ggrrrr/fibonacci-svc/common/api"
)

func TestSend(t *testing.T) {
	testWriter := http.TestResponseWriter{}
	send(&testWriter, api.Response{
		Payload: "data",
		Code:    200,
		Message: "msg",
	})
	assert.Equal(t, `{"payload":"data","code":200,"message":"msg"}`, testWriter.Output)
	assert.Equal(t, 200, testWriter.StatusCode)
}

func TestSendError(t *testing.T) {
	testWriter := http.TestResponseWriter{}
	SendError(&testWriter, errors.New("error msg"))
	assert.Equal(t, `{"code":500,"message":"error msg"}`, testWriter.Output)
	assert.Equal(t, 500, testWriter.StatusCode)
}

func TestSendPayload(t *testing.T) {
	testWriter := http.TestResponseWriter{}
	SendPayload(&testWriter, "data")
	assert.Equal(t, `{"payload":"data","code":200}`, testWriter.Output)
	assert.Equal(t, 200, testWriter.StatusCode)
}
