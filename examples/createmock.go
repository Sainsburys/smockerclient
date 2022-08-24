package main

import (
	"log"
	"net/http"

	"github.com/churmd/smockerclient"
	"github.com/churmd/smockerclient/mock"
)

func main() {
	instance := smockerclient.DefaultInstance()

	// Clear any old sessions and mocks
	err := instance.ResetAllSessionsAndMocks()
	if err != nil {
		log.Fatal(err)
	}

	// Start a new session for your new mocks
	err = instance.StartSession("SmockerClientSession")
	if err != nil {
		log.Fatal(err)
	}

	// Add a healthcheck mock
	requestBuilder := mock.NewRequestBuilder(http.MethodGet, "/healthcheck")
	requestBuilder.AddHeader("Accept", "application/json")
	request := requestBuilder.Build()

	responseBuilder := mock.NewResponseBuilder(http.StatusOK)
	responseBuilder.AddBody(`{"status": "OK"}`)
	response := responseBuilder.Build()

	mockDefinition := mock.NewDefinition(request, response)

	err = instance.AddMock(mockDefinition)
	if err != nil {
		log.Fatal(err)
	}
}
