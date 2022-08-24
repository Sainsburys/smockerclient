package mock_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

func TestNewResponseBuilder_defaultsTo200Response(t *testing.T) {
	expectedJson := `{"status": 200}`

	responseBuilder := mock.NewResponseBuilder()

	jsonBytes, err := responseBuilder.ToResponseJson()

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}

func TestNewNoContentResponse_ToResponseJson_CreatesResponseWithStatus204(t *testing.T) {
	expectedJson := `{"status": 204}`

	responseBuilder := mock.NewNoContentResponse()

	jsonBytes, err := responseBuilder.ToResponseJson()

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}
