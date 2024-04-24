package mock_test

import (
	"github.com/churmd/smockerclient/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMockCallLimit(t *testing.T) {
	expectedContext := &mock.Context{Times: 3}

	actualContext := mock.NewMockCallLimit(3)

	assert.Equal(t, expectedContext, actualContext)
}
