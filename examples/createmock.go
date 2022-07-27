package main

import (
	"log"

	"github.com/churmd/smockerclient"
	"github.com/churmd/smockerclient/mock"
)

func main() {
	instance := smockerclient.DefaultInstance()

	err := instance.StartSession("SmockerClientSession")
	if err != nil {
		log.Fatal(err)
	}

	exampleMockDefinition := `{
	   "request": {
		  "method": "GET",
		  "path": "/example"
	   },
	   "response": {
		  "status": 200,
		  "body": "{\"status\": \"OK\"}"
	   }
	}`
	jsonMock := mock.NewRawJsonDefinition(exampleMockDefinition)
	err = instance.AddMock(jsonMock)
	if err != nil {
		log.Fatal(err)
	}
}
