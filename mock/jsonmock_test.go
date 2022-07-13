package mock_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

func TestJsonMockReturnsGivenJson(t *testing.T) {
	json := jsonForMock()

	jsonMock := mock.NewJsonMock(json)
	definition, err := jsonMock.MarshalJSON()

	assert.NoError(t, err)
	assert.JSONEq(t, json, string(definition))
}

func jsonForMock() string {
	return `{
	   "request": {
		  "method": "GET",
		  "path": "/example"
	   },
	   "response": {
		  "status": 200,
		  "body": "{\"status\": \"OK\"}"
	   }
	}`
}
