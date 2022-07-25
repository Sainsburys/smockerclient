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
	req, err := i.createSessionRequest(name)
	if err != nil {
		return fmt.Errorf("smockerclient unable to create request to start a new session %w", err)
	}

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("smockerclient unable to send request to start a new session %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("smockerclient unable to create a new session named %s received status:%d", name, resp.StatusCode)
		}
		return fmt.Errorf("smockerclient unable to create a new session named %s received status:%d and message:%s", name, resp.StatusCode, body)
	}

	return nil
}

func (i Instance) createSessionRequest(name string) (*http.Request, error) {
	url := i.url + "/sessions"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("name", name)
	req.URL.RawQuery = query.Encode()

	return req, nil
}

func (i Instance) AddMock(mock Mock) error {
	req, err := i.createAddMockRequest(mock)
	if err != nil {
		return err
	}

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("smockerclient unable to send request to add a new mock %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("smockerclient unable to add mock and unable to read response message received status:%d", resp.StatusCode)
		}
		return fmt.Errorf("smockerclient unable to add mock received status:%d and message:%s", resp.StatusCode, body)
	}

	return nil
}

func (i Instance) createAddMockRequest(mock Mock) (*http.Request, error) {
	url := i.url + "/mocks"

	// Smocker API always expects a list of mocks to be sent
	mocks := []Mock{mock}
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(mocks)
	if err != nil {
		return nil, fmt.Errorf("smockerclient unable to create request body from mock %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("smockerclient unable to create request to add a new mock %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	return req, nil
}
