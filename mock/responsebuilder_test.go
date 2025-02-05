package mock_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Sainsburys/smockerclient/mock"
)

func TestResponseBuilder_Build(t *testing.T) {
	expectedResponse := mock.Response{
		Status: http.StatusOK,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
		Body: "{\"status\": \"OK\"}",
	}

	response := mock.NewResponseBuilder(http.StatusOK).
		AddBody(`{"status": "OK"}`).
		AddHeader("Content-Type", "application/json").
		Build()

	assert.Equal(t, expectedResponse, response)
}

func TestNewResponseBuilder_Build(t *testing.T) {
	expectedResponse := mock.Response{
		Status: http.StatusOK,
	}

	responseBuilder := mock.NewResponseBuilder(http.StatusOK)
	response := responseBuilder.Build()

	assert.Equal(t, expectedResponse, response)
}

func TestNewResponseBuilder_AddHeader_Build(t *testing.T) {
	expectedResponse := mock.Response{
		Status: http.StatusOK,
		Headers: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}

	responseBuilder := mock.NewResponseBuilder(http.StatusOK)
	responseBuilder.AddHeader("Content-Type", "application/json")
	response := responseBuilder.Build()

	assert.Equal(t, expectedResponse, response)
}

func TestNewResponseBuilder_AddBody(t *testing.T) {
	expectedResponse := mock.Response{
		Status: http.StatusOK,
		Body:   "{\"status\": \"OK\"}",
	}

	responseBuilder := mock.NewResponseBuilder(http.StatusOK)
	responseBuilder.AddBody(`{"status": "OK"}`)
	response := responseBuilder.Build()

	assert.Equal(t, expectedResponse, response)
}
