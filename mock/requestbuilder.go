package mock

import (
	"encoding/base64"
	"encoding/json"
)

type RequestBuilder struct {
	Method      string              `json:"method"`
	Path        string              `json:"path"`
	QueryParams map[string][]string `json:"query_params,omitempty"`
	Headers     map[string][]string `json:"headers,omitempty"`
	Body        *RequestBody        `json:"body,omitempty"`
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
	rb.initialiseQueryParams()
	rb.QueryParams[key] = values
}

func (rb *RequestBuilder) initialiseQueryParams() {
	if rb.QueryParams == nil {
		rb.QueryParams = make(map[string][]string, 1)
	}
}

func (rb *RequestBuilder) AddBearerAuthToken(token string) {
	bearerToken := "Bearer " + token
	rb.AddHeader("Authorization", bearerToken)
}

func (rb *RequestBuilder) AddBasicAuth(username string, password string) {
	basicToken := createBasicToken(username, password)
	rb.AddHeader("Authorization", basicToken)
}

func createBasicToken(username string, password string) string {
	usernamePasswordCombined := username + ":" + password
	base64Encoding := base64.StdEncoding.EncodeToString([]byte(usernamePasswordCombined))
	basicToken := "Basic " + base64Encoding
	return basicToken
}

func (rb *RequestBuilder) AddHeader(key string, values ...string) {
	rb.initialiseHeaders()
	rb.Headers[key] = values
}

func (rb *RequestBuilder) initialiseHeaders() {
	if rb.Headers == nil {
		rb.Headers = make(map[string][]string, 1)
	}
}

func (rb *RequestBuilder) AddJsonBody(jsonBody string) {
	body := NewJsonRequestBody(jsonBody)
	rb.Body = &body
}
