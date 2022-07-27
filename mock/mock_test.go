package mock_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

func TestMockJsonEncoding(t *testing.T) {
	reqJson := `{
		"method": "PUT",
		"path": "/foo/bar"
	}`
	fakeReq := FakeRequest{Json: reqJson}

	respJson := `{
		"status": 200,
		"body": "{\"data\": [\"foo\", \"bar\"]}"
	}`
	fakeResp := FakeResponse{Json: respJson}

	expectedJson := `{
		"request": {
			"method": "PUT",
			"path": "/foo/bar"
		},
		"response": {
			"status": 200,
			"body": "{\"data\": [\"foo\", \"bar\"]}"
		}
	}`

	definition := mock.Definition{
		Request:  fakeReq,
		Response: fakeResp,
	}

	actualJson, err := json.Marshal(definition)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(actualJson))
}

type FakeRequest struct {
	Json string
}

func (fr FakeRequest) MarshalJSON() ([]byte, error) {
	return []byte(fr.Json), nil
}

type FakeResponse struct {
	Json string
}

func (fr FakeResponse) MarshalJSON() ([]byte, error) {
	return []byte(fr.Json), nil
}
