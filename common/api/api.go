package api

type (
	// We use this model as generic response for all HTTP calls to this service.
	// including errors
	Response struct {
		Payload any    `json:"payload,omitempty"`
		Code    int    `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}
)
