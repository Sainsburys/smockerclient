package mock_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Sainsburys/smockerclient/mock"
)

func TestRequestBuilder_Build(t *testing.T) {
	expectedQueryParams := map[string][]string{
		"limit":   {"10"},
		"filters": {"red", "green"},
	}
	expectedHeaders := map[string][]string{
		"Content-Type":  {"application/json", "application/vnd.api+json"},
		"Authorization": {"Bearer sv2361fr1o8ph3oin"},
	}
	expectedReqBody := mock.RequestBody{
		Matcher: "ShouldEqualJSON",
		Value:   "{\"name\": \"John Smith\", \"uuid\": \"daa7b90d-9429-4d7a-9304-edc41ff44a6d\", \"rank\": 10}",
	}
	expectedRequest := mock.Request{
		Method:      http.MethodPut,
		Path:        "/foo/bar",
		QueryParams: expectedQueryParams,
		Headers:     expectedHeaders,
		Body:        &expectedReqBody,
	}

	jsonBody := `{"name": "John Smith", "uuid": "daa7b90d-9429-4d7a-9304-edc41ff44a6d", "rank": 10}`

	request := mock.NewRequestBuilder(http.MethodPut, "/foo/bar").
		AddQueryParam("limit", "10").
		AddQueryParam("filters", "red", "green").
		AddHeader("Content-Type", "application/json", "application/vnd.api+json").
		AddHeader("Authorization", "Bearer sv2361fr1o8ph3oin").
		AddJsonBody(jsonBody).
		Build()

	assert.Equal(t, expectedRequest, request)
}

func TestNewRequestBuilder_Build(t *testing.T) {
	expectedRequest := mock.Request{
		Method: http.MethodPut,
		Path:   "/foo/bar",
	}

	request := mock.NewRequestBuilder(http.MethodPut, "/foo/bar").Build()

	assert.Equal(t, expectedRequest, request)
}

func TestNewRequestBuilder_AddQueryParam_Build(t *testing.T) {
	expectedQueryParams := map[string][]string{
		"limit":   {"10"},
		"filters": {"red", "green"},
	}
	expectedRequest := mock.Request{
		Method:      http.MethodPut,
		Path:        "/foo/bar",
		QueryParams: expectedQueryParams,
	}

	requestBuilder := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	requestBuilder.AddQueryParam("limit", "10")
	requestBuilder.AddQueryParam("filters", "red", "green")
	request := requestBuilder.Build()

	assert.Equal(t, expectedRequest, request)
}

func TestNewRequestBuilder_AddHeader_Build(t *testing.T) {
	expectedHeaders := map[string][]string{
		"Content-Type":            {"application/json", "application/vnd.api+json"},
		"Authorization":           {"Bearer sv2361fr1o8ph3oin"},
		"Canonical-Header-Format": {"some-value"},
	}
	expectedRequest := mock.Request{
		Method:  http.MethodPut,
		Path:    "/foo/bar",
		Headers: expectedHeaders,
	}

	requestBuilder := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	requestBuilder.AddHeader("Content-Type", "application/json", "application/vnd.api+json")
	requestBuilder.AddHeader("Authorization", "Bearer sv2361fr1o8ph3oin")
	requestBuilder.AddHeader("canonical-header-FORMAT", "some-value")
	request := requestBuilder.Build()

	assert.Equal(t, expectedRequest, request)
}

func TestNewRequestBuilder_AddJsonBody_Build(t *testing.T) {
	expectedReqBody := mock.RequestBody{
		Matcher: "ShouldEqualJSON",
		Value:   "{\"name\": \"John Smith\", \"uuid\": \"daa7b90d-9429-4d7a-9304-edc41ff44a6d\", \"rank\": 10}",
	}
	expectedRequest := mock.Request{
		Method: http.MethodPut,
		Path:   "/foo/bar",
		Body:   &expectedReqBody,
	}

	jsonBody := `{"name": "John Smith", "uuid": "daa7b90d-9429-4d7a-9304-edc41ff44a6d", "rank": 10}`

	requestBuilder := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	requestBuilder.AddJsonBody(jsonBody)
	request := requestBuilder.Build()

	assert.Equal(t, expectedRequest, request)
}

func TestNewRequestBuilder_AddBearerAuthToken(t *testing.T) {
	expectedHeaders := map[string][]string{
		"Authorization": {"Bearer sv2361fr1o8ph3oin"},
	}
	expectedRequest := mock.Request{
		Method:  http.MethodPut,
		Path:    "/foo/bar",
		Headers: expectedHeaders,
	}

	requestBuilder := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	requestBuilder.AddBearerAuthToken("sv2361fr1o8ph3oin")
	request := requestBuilder.Build()

	assert.Equal(t, expectedRequest, request)
}

func TestNewRequestBuilder_AddBasicAuth(t *testing.T) {
	username := "admin"
	password := "password"

	expectedHeaders := map[string][]string{
		"Authorization": {"Basic YWRtaW46cGFzc3dvcmQ="},
	}
	expectedRequest := mock.Request{
		Method:  http.MethodPut,
		Path:    "/foo/bar",
		Headers: expectedHeaders,
	}

	requestBuilder := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	requestBuilder.AddBasicAuth(username, password)
	request := requestBuilder.Build()

	assert.Equal(t, expectedRequest, request)
}
