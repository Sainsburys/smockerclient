package mock_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient/mock"
)

func TestMultiMapJsonEncoding(t *testing.T) {
	expectedJson := `{
		"limit": ["10"],
		"key": ["value"],
		"filter": ["10", "20"]
	}`

	queryParams := mock.MultiMap{
		"limit":  []string{"10"},
		"key":    []string{"value"},
		"filter": []string{"10", "20"},
	}
	jsonBytes, err := json.Marshal(queryParams)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedJson, string(jsonBytes))
}
