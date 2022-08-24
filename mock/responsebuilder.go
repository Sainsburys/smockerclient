package mock

import (
	"encoding/json"
)

type ResponseBuilder struct {
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers,omitempty"`
	Body    string              `json:"body,omitempty"`
}

func NewResponseBuilder(httpStatus int) ResponseBuilder {
	return ResponseBuilder{
		Status: httpStatus,
	}
}

func (rb ResponseBuilder) ToResponseJson() ([]byte, error) {
	return json.Marshal(rb)
}

func (rb *ResponseBuilder) AddHeader(key string, values ...string) {
	rb.initialiseHeaders()
	rb.Headers[key] = values
}

func (rb *ResponseBuilder) initialiseHeaders() {
	if rb.Headers == nil {
		rb.Headers = make(map[string][]string, 1)
	}
}

func (rb *ResponseBuilder) AddBody(body string) {
	rb.Body = body
}
