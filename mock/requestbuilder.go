package mock

import (
	"encoding/json"
)

type RequestBuilder struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	QueryParams MultiMap `json:"query_params,omitempty"`
	Headers     MultiMap `json:"headers,omitempty"`
	Body        *Body    `json:"body,omitempty"`
}

func NewRequestBuilder(method, path string) RequestBuilder {
	return RequestBuilder{
		Method: method,
		Path:   path,
	}
}

func (rb RequestBuilder) ToRequestJson() ([]byte, error) {
	return json.Marshal(rb)
}

func (rb *RequestBuilder) AddQueryParam(key string, values ...string) {
	if rb.QueryParams == nil {
		rb.QueryParams = MultiMap{}
	}

	rb.QueryParams[key] = values
}

func (rb *RequestBuilder) AddHeader(key string, values ...string) {
	if rb.Headers == nil {
		rb.Headers = MultiMap{}
	}

	rb.Headers[key] = values
}

func (rb *RequestBuilder) AddJsonBody(jsonBody string) {
	body := NewJsonBody(jsonBody)
	rb.Body = &body
}
