package smockerclient_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/churmd/smockerclient"
	"github.com/stretchr/testify/assert"
)

func TestStartSession(t *testing.T) {
	serverCallCount := 0
	sessionName := "my-new-session"

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				serverCallCount++

				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/sessions", r.URL.Path)
				actualSessionName := r.URL.Query().Get("name")
				assert.Equal(t, sessionName, actualSessionName)

				resp := `{
					"id": "1d6d264b-4d13-4e0b-a51e-e44fc80eca9f",
					"name": "` + sessionName + `"
				  }`
				w.Write([]byte(resp))
			},
		),
	)
	defer server.Close()

	smockerInstance := smockerclient.NewInstance(server.URL)
	err := smockerInstance.StartSession(sessionName)

	assert.NoError(t, err)
	assert.Equal(t, 1, serverCallCount)
}

func TestAddMock(t *testing.T) {
	serverCallCount := 0
	expectedJson := `{"example": 1234}`
	fakeMock := FakeMock{Json: expectedJson}

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				serverCallCount++

				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/mocks", r.URL.Path)

				body, err := ioutil.ReadAll(r.Body)
				assert.NoError(t, err)
				assert.JSONEq(t, expectedJson, string(body))

				resp := `{
					"message": "Mocks registered successfully"
				  }`
				w.Write([]byte(resp))
			},
		),
	)
	defer server.Close()

	smockerInstance := smockerclient.NewInstance(server.URL)
	err := smockerInstance.AddMock(fakeMock)

	assert.NoError(t, err)
	assert.Equal(t, 1, serverCallCount)
}

type FakeMock struct {
	Json string
}

func (fm FakeMock) ToJsonDefinition() []byte {
	return []byte(fm.Json)
}
