package smockerclient

import (
	"bytes"
	"fmt"
	"net/http"
)

type Mock interface {
	ToJson() []byte
}

type Instance struct {
	url        string
	httpClient *http.Client
}

func DefaultInstance() Instance {
	return Instance{
		url:        "http://localhost:8081",
		httpClient: http.DefaultClient,
	}
}

func NewInstance(url string) Instance {
	return Instance{
		url:        url,
		httpClient: http.DefaultClient,
	}
}

func (i Instance) StartSession(name string) error {
	url := i.url + "/sessions?name=" + name
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("smockerclient unable to create request to start a new session %w", err)
	}

	_, err = i.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("smockerclient unable to send request to start a new session %w", err)
	}

	return nil
}

func (i Instance) AddMock(mock Mock) error {
	url := i.url + "/mocks"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(mock.ToJson()))
	if err != nil {
		return fmt.Errorf("smockerclient unable to create request to add a new mock %w", err)
	}

	_, err = i.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("smockerclient unable to send request to add a new mock %w", err)
	}

	return nil
}
