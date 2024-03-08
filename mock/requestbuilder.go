package mock

import (
	"encoding/base64"
	"net/http"
)

type RequestBuilder struct {
	request *Request
}

func NewRequestBuilder(method, path string) RequestBuilder {
	request := Request{
		Method: method,
		Path:   path,
	}
	return RequestBuilder{
		request: &request,
	}
}

func (rb RequestBuilder) Build() Request {
	return *rb.request
}

func (rb RequestBuilder) AddQueryParam(key string, values ...string) RequestBuilder {
	rb.initialiseQueryParams()
	rb.request.QueryParams[key] = values

	return rb
}

func (rb RequestBuilder) initialiseQueryParams() {
	if rb.request.QueryParams == nil {
		rb.request.QueryParams = make(map[string][]string, 1)
	}
}

func (rb RequestBuilder) AddBearerAuthToken(token string) RequestBuilder {
	bearerToken := "Bearer " + token
	rb.AddHeader("Authorization", bearerToken)

	return rb
}

func (rb RequestBuilder) AddBasicAuth(username string, password string) RequestBuilder {
	basicToken := createBasicToken(username, password)
	rb.AddHeader("Authorization", basicToken)

	return rb
}

func createBasicToken(username string, password string) string {
	usernamePasswordCombined := username + ":" + password
	base64Encoding := base64.StdEncoding.EncodeToString([]byte(usernamePasswordCombined))
	basicToken := "Basic " + base64Encoding

	return basicToken
}

// AddHeader Sets a request header key and values to match against on the mock definition
//
// Note: the header key will be formatted to [http.CanonicalHeaderKey] as smocker is case-sensitive regarding header
// keys, this will help prevent mock expectations not matching as Go's standard http library uses the Canonical format.
func (rb RequestBuilder) AddHeader(key string, values ...string) RequestBuilder {
	rb.initialiseHeaders()
	canonicalHeaderKey := http.CanonicalHeaderKey(key)
	rb.request.Headers[canonicalHeaderKey] = values

	return rb
}

func (rb RequestBuilder) initialiseHeaders() {
	if rb.request.Headers == nil {
		rb.request.Headers = make(map[string][]string, 1)
	}
}

func (rb RequestBuilder) AddJsonBody(jsonBody string) RequestBuilder {
	body := RequestBody{
		Matcher: "ShouldEqualJSON",
		Value:   jsonBody,
	}
	rb.request.Body = &body

	return rb
}
