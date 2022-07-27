package mock_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

func TestRequestBuilder_ToRequestJson(t *testing.T) {
	expectedJson := `{
		"method": "PUT",
		"path": "/foo/bar",
		"query_params": {
            "limit": ["10"],
            "filters": ["red", "green"]
		},
		"headers": {
			"Content-Type": ["application/json", "application/vnd.api+json"],
			"Authorization": ["Bearer sv2361fr1o8ph3oin"]
		},
		"body": {
			"matcher": "ShouldEqualJSON",
			"value": "{\"name\":\"John Smith\",\"uuid\":\"daa7b90d-9429-4d7a-9304-edc41ff44a6d\",\"rank\":10}"
		}
	}`

	request := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	request.AddQueryParam("limit", "10")
	request.AddQueryParam("filters", "red", "green")
	request.AddHeader("Content-Type", "application/json", "application/vnd.api+json")
	request.AddHeader("Authorization", "Bearer sv2361fr1o8ph3oin")

	jsonBody := `{
		"name": "John Smith",
		"uuid": "daa7b90d-9429-4d7a-9304-edc41ff44a6d",
		"rank": 10
	}`
	err := request.AddJsonBody(jsonBody)
	assert.NoError(t, err)

	jsonBytes, err := json.Marshal(request)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestRequestBuilderBuilderBasicJsonEncoding(t *testing.T) {
	expectedJson := `{
		"method": "PUT",
		"path": "/foo/bar"
	}`

	request := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")

	jsonBytes, err := json.Marshal(request)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestRequestBuilderWithQueryParamsJsonEncoding(t *testing.T) {
	expectedJson := `{
		"method": "PUT",
		"path": "/foo/bar",
		"query_params": {
            "limit": ["10"],
            "filters": ["red", "green"]
		}
	}`

	request := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	request.AddQueryParam("limit", "10")
	request.AddQueryParam("filters", "red", "green")

	jsonBytes, err := json.Marshal(request)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestRequestBuilderWithHeadersJsonEncoding(t *testing.T) {
	expectedJson := `{
		"method": "PUT",
		"path": "/foo/bar",
		"headers": {
			"Content-Type": ["application/json", "application/vnd.api+json"],
			"Authorization": ["Bearer sv2361fr1o8ph3oin"]
		}
	}`

	request := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	request.AddHeader("Content-Type", "application/json", "application/vnd.api+json")
	request.AddHeader("Authorization", "Bearer sv2361fr1o8ph3oin")

	jsonBytes, err := json.Marshal(request)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestRequestBuilderWithJsonBodyEncoding(t *testing.T) {
	jsonBody := `{
		"name": "John Smith",
		"uuid": "daa7b90d-9429-4d7a-9304-edc41ff44a6d",
		"rank": 10
	}`
	expectedJson := `{
		"method": "PUT",
		"path": "/foo/bar",
		"body": {
			"matcher": "ShouldEqualJSON",
			"value": "{\"name\":\"John Smith\",\"uuid\":\"daa7b90d-9429-4d7a-9304-edc41ff44a6d\",\"rank\":10}"
		}
	}`

	request := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	err := request.AddJsonBody(jsonBody)

	assert.NoError(t, err)

	jsonBytes, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestRequestBuilderWithJsonBodyGivenBadJsonErrors(t *testing.T) {
	jsonBody := `{name: "example"}`

	request := mock.NewRequestBuilder(http.MethodPut, "/foo/bar")
	err := request.AddJsonBody(jsonBody)

	assert.ErrorContains(t, err, "unable to compact body json")
}

func TestMultiMapJsonEncoding(t *testing.T) {
	expectedJson := `{
		"limit": ["10"],
		"key": ["value"],
		"filter": ["10", "20"]
	}`

	queryParams := mock.MultiMap{
		"limit":  []string{"10"},
		"key":    []string{"value"},
		"filter": []string{"10", "20"},
	}
	jsonBytes, err := json.Marshal(queryParams)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}
