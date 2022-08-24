package smockerclient_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/churmd/smockerclient"
)

const jsonContentType = "application/json"

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

func TestStartSession_WithNameThatNeedsUrlEscaping(t *testing.T) {
	serverCallCount := 0
	sessionName := `test !@£$%^&*()`

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

func TestStartSession_WhenServerDoesNotReturn200_ReturnsError(t *testing.T) {
	sessionName := "my-new-session"

	server, serverCallCount := newBadResponseServer()
	defer server.Close()

	smockerInstance := smockerclient.NewInstance(server.URL)
	err := smockerInstance.StartSession(sessionName)

	assert.Equal(t, 1, *serverCallCount)
	assert.EqualError(t, err, "smockerclient unable to create a new session named my-new-session. received status:400 and message:400 Bad Request")
}

func TestAddMock(t *testing.T) {
	serverCallCount := 0
	jsonMock := `{"example": 1234}`
	fakeMock := FakeMock{Json: jsonMock}
	expectedJson := "[" + jsonMock + "]"

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				serverCallCount++

				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/mocks", r.URL.Path)

				contentType := r.Header.Get("Content-Type")
				assert.Equal(t, jsonContentType, contentType)

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

func TestAddMock_WhenMockJsonConversionErrors_ReturnsError(t *testing.T) {
	mockError := errors.New("fails mock json conversion")
	fakeMock := FakeMock{Error: mockError}

	serverCallCount := 0
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				serverCallCount++
			},
		),
	)
	defer server.Close()

	smockerInstance := smockerclient.NewInstance(server.URL)
	err := smockerInstance.AddMock(fakeMock)

	assert.Equal(t, 0, serverCallCount)
	expectedError := fmt.Errorf("unable to convert mock to json when running ToMockDefinitionJson. %w", mockError)
	assert.ErrorContains(t, err, expectedError.Error())
}

func TestAddMock_WhenServerDoesNotReturn200_ReturnsError(t *testing.T) {
	expectedJson := `{"example": 1234}`
	fakeMock := FakeMock{Json: expectedJson}

	server, serverCallCount := newBadResponseServer()
	defer server.Close()

	smockerInstance := smockerclient.NewInstance(server.URL)
	err := smockerInstance.AddMock(fakeMock)

	assert.Equal(t, 1, *serverCallCount)
	assert.EqualError(t, err, "smockerclient unable to add mock. received status:400 and message:400 Bad Request")
}

func TestResetAllSessionsAndMocks(t *testing.T) {
	serverCallCount := 0

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				serverCallCount++

				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/reset", r.URL.Path)

				resp := `{
			  		"message": "Reset successful"
				}`
				w.Write([]byte(resp))
			},
		),
	)
	defer server.Close()

	smockerInstance := smockerclient.NewInstance(server.URL)
	err := smockerInstance.ResetAllSessionsAndMocks()

	assert.NoError(t, err)
	assert.Equal(t, 1, serverCallCount)
}

func TestResetAllSessionsAndMocks_WhenServerDoesNotReturn200_ReturnsError(t *testing.T) {
	server, serverCallCount := newBadResponseServer()
	defer server.Close()

	smockerInstance := smockerclient.NewInstance(server.URL)
	err := smockerInstance.ResetAllSessionsAndMocks()

	assert.Equal(t, 1, *serverCallCount)
	assert.EqualError(t, err, "smockerclient unable to reset all the sessions and mocks. received status:400 and message:400 Bad Request")
}

func newBadResponseServer() (*httptest.Server, *int) {
	serverCallCount := 0

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				serverCallCount++
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 Bad Request"))
			},
		),
	)

	return server, &serverCallCount
}

type FakeMock struct {
	Json  string
	Error error
}

func (fm FakeMock) ToMockDefinitionJson() ([]byte, error) {
	return []byte(fm.Json), fm.Error
}
