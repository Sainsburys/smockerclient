package mock_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

func TestDefinition_ToMockJson(t *testing.T) {
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

	definition := mock.NewDefinition(fakeReq, fakeResp)

	actualJson, err := definition.ToMockJson()

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(actualJson))
}

type FakeRequest struct {
	Json string
}

func (fr FakeRequest) ToRequestJson() ([]byte, error) {
	return []byte(fr.Json), nil
}

type FakeResponse struct {
	Json string
}

func (fr FakeResponse) ToResponseJson() ([]byte, error) {
	return []byte(fr.Json), nil
}
