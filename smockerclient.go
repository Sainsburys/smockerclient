package smockerclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Mock json.Marshaler

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
	// TODO handle query encoding
	url := i.url + "/sessions?name=" + name
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("smockerclient unable to create request to start a new session %w", err)
	}

	_, err = i.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("smockerclient unable to send request to start a new session %w", err)
	}

	// TODO check status code

	return nil
}

func (i Instance) AddMock(mock Mock) error {
	url := i.url + "/mocks"

	// Smocker API always expects a list of mocks to be sent
	mocks := []Mock{mock}
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(mocks)
	if err != nil {
		return fmt.Errorf("smockerclient unable to create request body from mock %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("smockerclient unable to create request to add a new mock %w", err)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("smockerclient unable to send request to add a new mock %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unable to add mock and unable to read response message received status:%d", resp.StatusCode)
		}
		return fmt.Errorf("unable to add mock received status:%d and message:%s", resp.StatusCode, body)
	}

	return nil
}
