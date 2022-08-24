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
	request := mock.NewRequestBuilder(http.MethodGet, "/healthcheck")
	request.AddHeader("Accept", "application/json")

	response := mock.NewResponseBuilder(http.StatusOK)
	response.AddBody(`{"status": "OK"}`)

	mockDefinition := mock.NewDefinition(request, response)

	err = instance.AddMock(mockDefinition)
	if err != nil {
		log.Fatal(err)
	}
}
