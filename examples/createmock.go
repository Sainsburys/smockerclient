package main

import (
	"log"
	"net/http"

	"github.com/Sainsburys/smockerclient"
	"github.com/Sainsburys/smockerclient/mock"
)

func main() {
	instance := smockerclient.Instance{}

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
	request := mock.NewRequestBuilder(http.MethodGet, "/healthcheck").
		AddHeader("Accept", "application/json").
		Build()

	response := mock.NewResponseBuilder(http.StatusOK).AddBody(`{"status": "OK"}`).Build()

	mockDefinition := mock.NewDefinition(request, response)

	err = instance.AddMock(mockDefinition)
	if err != nil {
		log.Fatal(err)
	}

	// Call the healthcheck mock in the code under test
	someCodeUnderTest()

	// Verify all the mocks were used and no extra requests were made
	err = instance.VerifyMocksInCurrentSession()
	if err != nil {
		log.Fatal(err)
	}
}

func someCodeUnderTest() {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/healthcheck", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Accept", "application/json")

	_, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
}
