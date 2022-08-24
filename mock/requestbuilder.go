package mock

import (
	"encoding/base64"
)

type RequestBuilder struct {
	request Request
}

func NewRequestBuilder(method, path string) RequestBuilder {
	return RequestBuilder{
		request: Request{
			Method: method,
			Path:   path,
		},
	}
}

func (rb RequestBuilder) Build() Request {
	return rb.request
}

func (rb *RequestBuilder) AddQueryParam(key string, values ...string) {
	rb.initialiseQueryParams()
	rb.request.QueryParams[key] = values
}

func (rb *RequestBuilder) initialiseQueryParams() {
	if rb.request.QueryParams == nil {
		rb.request.QueryParams = make(map[string][]string, 1)
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
	rb.request.Headers[key] = values
}

func (rb *RequestBuilder) initialiseHeaders() {
	if rb.request.Headers == nil {
		rb.request.Headers = make(map[string][]string, 1)
	}
}

func (rb *RequestBuilder) AddJsonBody(jsonBody string) {
	body := RequestBody{
		Matcher: "ShouldEqualJSON",
		Value:   jsonBody,
	}
	rb.request.Body = &body
}
