package mock_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

type TestRequestBody struct {
	Matcher string `json:"matcher"`
	Value   string `json:"value"`
}

func TestNewJsonBody_JsonEncoding(t *testing.T) {
	jsonBody := `{
		"name": "John Smith",
		"uuid": "daa7b90d-9429-4d7a-9304-edc41ff44a6d",
		"rank": 10
	}`

	body := mock.NewJsonRequestBody(jsonBody)

	actualJson, err := json.Marshal(body)

	assert.NoError(t, err)

	var testReqBody TestRequestBody
	err = json.Unmarshal(actualJson, &testReqBody)
	assert.NoError(t, err)
	assert.Equal(t, "ShouldEqualJSON", testReqBody.Matcher)
	assert.JSONEq(t, jsonBody, testReqBody.Value)
}
