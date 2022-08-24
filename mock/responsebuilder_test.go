package mock_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

func TestResponseBuilder_ToResponseJson(t *testing.T) {
	expectedJson := `{
		"status": 200,
		"headers": {
			"Content-Type": ["application/json"]
		},
		"body": "{\"status\": \"OK\"}"
	}`

	responseBuilder := mock.NewResponseBuilder(http.StatusOK)
	responseBuilder.AddBody(`{"status": "OK"}`)
	responseBuilder.AddHeader("Content-Type", "application/json")

	jsonBytes, err := responseBuilder.ToResponseJson()

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestNewResponseBuilder(t *testing.T) {
	expectedJson := `{"status": 200}`

	responseBuilder := mock.NewResponseBuilder(http.StatusOK)

	jsonBytes, err := responseBuilder.ToResponseJson()

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestNewResponseBuilder_AddHeader(t *testing.T) {
	expectedJson := `{
		"status": 200,
		"headers": {
			"Content-Type": ["application/json", "application/vnd.api+json"]
		}
	}`

	responseBuilder := mock.NewResponseBuilder(http.StatusOK)
	responseBuilder.AddHeader("Content-Type", "application/json", "application/vnd.api+json")

	jsonBytes, err := responseBuilder.ToResponseJson()

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestNewResponseBuilder_AddBody(t *testing.T) {
	expectedJson := `{
		"status": 200,
		"body": "{\"status\": \"OK\"}"
	}`
	jsonBody := `{"status": "OK"}`

	responseBuilder := mock.NewResponseBuilder(http.StatusOK)
	responseBuilder.AddBody(jsonBody)

	jsonBytes, err := responseBuilder.ToResponseJson()

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}
