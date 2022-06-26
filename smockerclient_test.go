package smockerclient_test

import (
	"testing"

	"github.com/churmd/smockerclient"
	"github.com/stretchr/testify/assert"
)

func TestDefaultInstance(t *testing.T) {
	expectedUrl := "http://localhost:8081"

	smockerInstance := smockerclient.DefaultInstance()

	assert.Equal(t, expectedUrl, smockerInstance.GetURL())
}
