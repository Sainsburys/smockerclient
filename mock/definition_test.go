package mock_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

func TestDefinition_ToMockJson(t *testing.T) {
	expectedJson := `{
		"request": {
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
				"value": "{\"name\": \"John Smith\", \"uuid\": \"daa7b90d-9429-4d7a-9304-edc41ff44a6d\", \"rank\": 10}"
			}
		},
		"response": {
			"status": 200,
			"headers": {
				"Content-Type": ["application/json"]
			},
			"body": "{\"status\": \"OK\"}"
		}
	}`

	request := createRequest()
	response := createResponse()
	definition := mock.NewDefinition(request, response)

	actualJson, err := definition.ToMockDefinitionJson()

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(actualJson))
}

func createRequest() mock.Request {
	reqQueryParams := map[string][]string{
		"limit":   {"10"},
		"filters": {"red", "green"},
	}
	reqHeaders := map[string][]string{
		"Content-Type":  {"application/json", "application/vnd.api+json"},
		"Authorization": {"Bearer sv2361fr1o8ph3oin"},
	}
	reqBody := mock.RequestBody{
		Matcher: "ShouldEqualJSON",
		Value:   "{\"name\": \"John Smith\", \"uuid\": \"daa7b90d-9429-4d7a-9304-edc41ff44a6d\", \"rank\": 10}",
	}
	request := mock.Request{
		Method:      http.MethodPut,
		Path:        "/foo/bar",
		QueryParams: reqQueryParams,
		Headers:     reqHeaders,
		Body:        &reqBody,
	}
	return request
}

func createResponse() mock.Response {
	response := mock.Response{
		Status: http.StatusOK,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: "{\"status\": \"OK\"}",
	}
	return response
}
