package mock_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

func TestNewJsonBody_JsonEncoding(t *testing.T) {
	expectedJson := `{
		"matcher": "ShouldEqualJSON",
		"value": "{\"name\":\"John Smith\",\"uuid\":\"daa7b90d-9429-4d7a-9304-edc41ff44a6d\",\"rank\":10}"
	}`

	jsonBody := `{
		"name": "John Smith",
		"uuid": "daa7b90d-9429-4d7a-9304-edc41ff44a6d",
		"rank": 10
	}`

	body := mock.NewJsonBody(jsonBody)

	actualJson, err := json.Marshal(body)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(actualJson))
}

func TestNewJsonBody_JsonEncodingWithBadBodyErrors(t *testing.T) {
	jsonBody := `{wrong: "json"}`

	body := mock.NewJsonBody(jsonBody)

	_, err := json.Marshal(body)

	assert.ErrorContains(t, err, "unable to compact body json")
}
