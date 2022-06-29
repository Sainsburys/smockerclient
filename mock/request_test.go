package mock_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/churmd/smockerclient/mock"
	"github.com/stretchr/testify/assert"
)

func TestRequestBasicJsonEncoding(t *testing.T) {
	expectedJson := `{
		"method": "PUT",
		"path": "/foo/bar"
	}`

	request := mock.NewRequest(http.MethodPut, "/foo/bar")


	jsonBytes, err := json.Marshal(request)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestRequestWithQueryParamsJsonEncoding(t *testing.T) {
	expectedJson := `{
		"method": "PUT",
		"path": "/foo/bar",
		"query_params": {
            "limit": "10",
            "offset": "0"
		}
	}`

	request := mock.NewRequest(http.MethodPut, "/foo/bar")
	request.AddQueryParam("limit", "10")
	request.AddQueryParam("offset", "0")

	jsonBytes, err := json.Marshal(request)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestQueryParamsJsonEncoding(t *testing.T) {
	expectedJson := `{
		"limit": "10",
		"key": "value"
	}`

	queryParams := mock.QueryParams{"limit":"10", "key":"value"}
	jsonBytes, err := json.Marshal(queryParams)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}