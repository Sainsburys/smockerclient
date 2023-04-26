package smockerclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// MockDefinition Allows multiple styles of mock creation to be used and custom extension.
// ToMockDefinitionJson must return json conforming to the smocker mock definition
// https://smocker.dev/technical-documentation/mock-definition.html as bytes.
type MockDefinition interface {
	ToMockDefinitionJson() ([]byte, error)
}

type Instance struct {
	url        string
	httpClient *http.Client
}

// DefaultInstance Creates an instance that will connect to the default smocker server location, http://localhost:8081
func DefaultInstance() Instance {
	return Instance{
		url:        "http://localhost:8081",
		httpClient: http.DefaultClient,
	}
}

// NewInstance Creates an instance that will connect to the smocker server using the url provided.
func NewInstance(url string) Instance {
	return Instance{
		url:        url,
		httpClient: http.DefaultClient,
	}
}

// StartSession Starts a new session on the Smocker server with the given name. New mocks will be added to the latest
// session started.
func (i Instance) StartSession(name string) error {
	resp, err := i.sendStartSessionRequest(name)
	if err != nil {
		return fmt.Errorf("smockerclient unable to create a new session named %s. %w", name, err)
	}

	err = handleNon200Response(resp)
	if err != nil {
		return fmt.Errorf("smockerclient unable to create a new session named %s. %w", name, err)
	}

	return nil
}

func (i Instance) sendStartSessionRequest(name string) (*http.Response, error) {
	req, err := i.createSessionRequest(name)
	if err != nil {
		return nil, fmt.Errorf("unable to create request. %w", err)
	}

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send request. %w", err)
	}

	return resp, nil
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

// AddMock Adds a new mock to the latest session on the Smocker server.
func (i Instance) AddMock(mock MockDefinition) error {
	resp, err := i.sendAddMockRequest(mock)
	if err != nil {
		return fmt.Errorf("smockerclient unable to add a new mock. %w", err)
	}

	err = handleNon200Response(resp)
	if err != nil {
		return fmt.Errorf("smockerclient unable to add mock. %w", err)
	}

	return nil
}

func (i Instance) sendAddMockRequest(mock MockDefinition) (*http.Response, error) {
	req, err := i.createAddMockRequest(mock)
	if err != nil {
		return nil, fmt.Errorf("unable to create request. %w", err)
	}

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to send request. %w", err)
	}

	return resp, nil
}

func (i Instance) createAddMockRequest(mock MockDefinition) (*http.Request, error) {
	body, err := createAddMockRequestBody(mock)
	if err != nil {
		return nil, err
	}

	url := i.url + "/mocks"
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func createAddMockRequestBody(mock MockDefinition) (*bytes.Buffer, error) {
	mockJson, err := mock.ToMockDefinitionJson()
	if err != nil {
		return nil, fmt.Errorf("unable to convert mock to json when running ToMockDefinitionJson. %w", err)
	}

	// Smocker API always expects a list of mocks to be sent
	mocks := []json.RawMessage{mockJson}
	body := &bytes.Buffer{}
	err = json.NewEncoder(body).Encode(mocks)
	if err != nil {
		return nil, fmt.Errorf("unable to create request body bytes from mock. %w", err)
	}

	return body, nil
}

// ResetAllSessionsAndMocks Clears the Smocker server of all sessions and mocks. Leaving it in a clean state
func (i Instance) ResetAllSessionsAndMocks() error {
	resp, err := i.sendResetAllSessionsAndMocksRequest()
	if err != nil {
		return fmt.Errorf("smockerclient unable to reset all the sessions and mocks. %w", err)
	}

	err = handleNon200Response(resp)
	if err != nil {
		return fmt.Errorf("smockerclient unable to reset all the sessions and mocks. %w", err)
	}

	return nil
}

func (i Instance) sendResetAllSessionsAndMocksRequest() (*http.Response, error) {
	request, err := i.createResetAllSessionAndMocksRequest()
	if err != nil {
		return nil, fmt.Errorf("unable to create request. %w", err)
	}

	resp, err := i.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("unable to send request. %w", err)
	}

	return resp, nil
}

func (i Instance) createResetAllSessionAndMocksRequest() (*http.Request, error) {
	url := i.url + "/reset"
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (i Instance) VerifyMocksInLatestSession() error {

	url := i.url + "/sessions/verify"
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		// TODO
		return err
	}

	response, err := i.httpClient.Do(request)
	if err != nil {
		// TODO
		return err
	}

	var verifiedResp verifiedResponse
	err = json.NewDecoder(response.Body).Decode(&verifiedResp)
	if err != nil {
		// TODO
		return err
	}

	if !verifiedResp.Mocks.AllUsed {
		return errors.New("not all mocks were used")
	}

	return nil
}

func handleNon200Response(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response message received status:%d", resp.StatusCode)
	}

	return fmt.Errorf("received status:%d and message:%s", resp.StatusCode, body)

}

type verifiedResponse struct {
	Mocks verifiedResponseMocks `json:"mocks"`
}

type verifiedResponseMocks struct {
	Verified bool   `json:"verified"`
	AllUsed  bool   `json:"all_used"`
	Message  string `json:"message"`
}
