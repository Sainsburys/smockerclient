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

func TestVerifyMocksInLatestSession_WhenAllMocksHaveBeenCalled_ReturnsNoError(t *testing.T) {
	serverCallCount := 0

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				serverCallCount++

				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/sessions/verify", r.URL.Path)

				resp := `{
				  "mocks": {
					"verified": true,
					"all_used": true,
					"message": "All mocks match expectations"
				  },
				  "history": {
					"verified": true,
					"message": "History is clean"
				  }
				}`
				w.Write([]byte(resp))
			},
		),
	)
	defer server.Close()

	smockerInstance := smockerclient.NewInstance(server.URL)
	err := smockerInstance.VerifyMocksInLatestSession()

	assert.NoError(t, err)
	assert.Equal(t, 1, serverCallCount)
}

func TestVerifyMocksInLatestSession_WhenSomeMocksHaveNotBeenCalled_ReturnsError(t *testing.T) {
	serverCallCount := 0

	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				serverCallCount++

				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/sessions/verify", r.URL.Path)

				resp := `{
				  "mocks": {
					"verified": false,
					"all_used": false,
					"message": "Some mocks don't match expectations",
					"failures": [
					  {
						"request": {
						  "path": "/test",
						  "method": "GET"
						},
						"response": {
						  "body": "{\"message\": \"test2\"}\n",
						  "status": 200,
						  "headers": {
							"Content-Type": ["application/json"]
						  }
						},
						"context": {
						  "times": 1
						},
						"state": {
						  "id": "6ecbd8f8-e2a7-4119-9be6-ad5ec83c58b6",
						  "times_count": 2,
						  "creation_date": "2020-02-26T12:11:34.713399+01:00"
						}
					  },
					  {
						"request": {
						  "path": "/test",
						  "method": "GET"
						},
						"response": {
						  "body": "{\"message\": \"test\"}\n",
						  "status": 200,
						  "headers": {
							"Content-Type": ["application/json"]
						  }
						},
						"context": {
						  "times": 1
						},
						"state": {
						  "id": "30266b21-77c0-48e6-b27e-5aa02d7defd8",
						  "times_count": 2,
						  "creation_date": "2020-02-26T12:11:34.713396+01:00"
						}
					  }
					],
					"unused": [
					  {
						"request": {
						  "path": "/test",
						  "method": "GET"
						},
						"response": {
						  "status": 200
						},
						"context": {},
						"state": {
						  "id": "d9ce47d4-86b7-4cb5-b7e9-829063704cec",
						  "times_count": 0,
						  "creation_date": "2020-02-26T12:11:34.747289+01:00"
						}
					  }
					]
				  },
				  "history": {
					"verified": false,
					"message": "There are errors in the history",
					"failures": [
					  {
						"request": {
						  "path": "/test",
						  "method": "GET",
						  "body": "",
						  "headers": {
							"Accept-Encoding": ["gzip"],
							"Host": ["localhost:8080"],
							"User-Agent": ["Go-http-client/1.1"]
						  },
						  "date": "2020-02-26T12:11:34.737809+01:00"
						},
						"response": {
						  "status": 666,
						  "body": {
							"message": "Matching mock found but was exceeded",
							"nearest": [
							  {
								"context": {
								  "times": 1
								},
								"request": {
								  "method": "GET",
								  "path": "/test"
								},
								"response": {
								  "body": "{\"message\": \"test2\"}\n",
								  "headers": {
									"Content-Type": ["application/json"]
								  },
								  "status": 200
								},
								"state": {
								  "creation_date": "2020-02-26T12:11:34.713399+01:00",
								  "id": "6ecbd8f8-e2a7-4119-9be6-ad5ec83c58b6",
								  "times_count": 2
								}
							  },
							  {
								"context": {
								  "times": 1
								},
								"request": {
								  "method": "GET",
								  "path": "/test"
								},
								"response": {
								  "body": "{\"message\": \"test\"}\n",
								  "headers": {
									"Content-Type": ["application/json"]
								  },
								  "status": 200
								},
								"state": {
								  "creation_date": "2020-02-26T12:11:34.713396+01:00",
								  "id": "30266b21-77c0-48e6-b27e-5aa02d7defd8",
								  "times_count": 2
								}
							  }
							],
							"request": {
							  "body": "",
							  "date": "2020-02-26T12:11:34.737814+01:00",
							  "headers": {
								"Accept-Encoding": ["gzip"],
								"Host": ["localhost:8080"],
								"User-Agent": ["Go-http-client/1.1"]
							  },
							  "method": "GET",
							  "path": "/test"
							}
						  },
						  "headers": {
							"Content-Type": ["application/json; charset=UTF-8"]
						  },
						  "date": "2020-02-26T12:11:34.738625+01:00"
						}
					  }
					]
				  }
				}`
				w.Write([]byte(resp))
			},
		),
	)
	defer server.Close()

	smockerInstance := smockerclient.NewInstance(server.URL)
	err := smockerInstance.VerifyMocksInLatestSession()

	assert.Error(t, err)
	assert.Equal(t, 1, serverCallCount)
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
