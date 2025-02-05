package mock_test

import (
	"github.com/Sainsburys/smockerclient/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithCallLimit(t *testing.T) {
	t.Run("it creates a new context with the times key if one doesnt exist", func(t *testing.T) {
		req := createRequest()
		res := createResponse()
		def := mock.NewDefinition(req, res, mock.WithCallLimit(3))

		expectedDef := mock.Definition{
			Request:  req,
			Response: res,
			Context:  &mock.Context{Times: 3},
		}
		assert.Equal(t, expectedDef, def)
	})

	t.Run("if a context already exists, it uses that", func(t *testing.T) {
		req := createRequest()
		res := createResponse()
		def := mock.NewDefinition(req, res, mock.WithCallLimit(3), mock.WithCallLimit(5))

		expectedDef := mock.Definition{
			Request:  req,
			Response: res,
			Context:  &mock.Context{Times: 5},
		}
		assert.Equal(t, expectedDef, def)
	})
}
