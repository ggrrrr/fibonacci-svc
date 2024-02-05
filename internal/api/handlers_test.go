package api

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/http"
	gomock "go.uber.org/mock/gomock"

	"github.com/ggrrrr/fibonacci-svc/internal/fi"
)

func TestHandlers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockApp := NewMockApp(mockCtrl)
	testHandlers := &handlers{app: mockApp}

	type testCase struct {
		name     string
		testFunc func(testWriter *http.TestResponseWriter)
		expect   string
	}
	tests := []testCase{
		{
			name:   "ok handleNext",
			expect: `{"payload":1,"code":200}`,
			testFunc: func(testWriter *http.TestResponseWriter) {
				mockApp.EXPECT().Next().Times(1).Return(fi.Number(1), nil)
				testHandlers.handleNext(testWriter, nil)
				fmt.Printf("Output: %v\n", testWriter.Output)
			},
		},
		{
			name:   "ok handlePrevious",
			expect: `{"payload":1,"code":200}`,
			testFunc: func(testWriter *http.TestResponseWriter) {
				mockApp.EXPECT().Previous().Times(1).Return(fi.Number(1), nil)
				testHandlers.handlePrevious(testWriter, nil)
				fmt.Printf("Output: %v\n", testWriter.Output)
			},
		},
		{
			name:   "ok handleCurrent",
			expect: `{"payload":1,"code":200}`,
			testFunc: func(testWriter *http.TestResponseWriter) {
				mockApp.EXPECT().Current().Times(1).Return(fi.Number(1))
				testHandlers.handleCurrent(testWriter, nil)
				fmt.Printf("Output: %v\n", testWriter.Output)
			},
		},
		{
			name:   "error handleNext",
			expect: `{"code":500,"message":"error"}`,
			testFunc: func(testWriter *http.TestResponseWriter) {
				mockApp.EXPECT().Next().Times(1).Return(fi.Number(0), errors.New("error"))
				testHandlers.handleNext(testWriter, nil)
				fmt.Printf("Output: %v\n", testWriter.Output)
			},
		},
		{
			name:   "error Previous",
			expect: `{"code":500,"message":"error"}`,
			testFunc: func(testWriter *http.TestResponseWriter) {
				mockApp.EXPECT().Previous().Times(1).Return(fi.Number(0), errors.New("error"))
				testHandlers.handlePrevious(testWriter, nil)
				fmt.Printf("Output: %v\n", testWriter.Output)
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testWriter := &http.TestResponseWriter{}
			tc.testFunc(testWriter)
			assert.Equal(t, tc.expect, testWriter.Output)
		})
	}
}
