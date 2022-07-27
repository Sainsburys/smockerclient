package mock

import (
	"encoding/json"
)

type Request interface {
	ToRequestJson() ([]byte, error)
}

type Response interface {
	ToResponseJson() ([]byte, error)
}

type Definition struct {
	Request  Request
	Response Response
}

func NewDefinition(req Request, resp Response) Definition {
	return Definition{
		Request:  req,
		Response: resp,
	}
}

func (d Definition) ToMockJson() ([]byte, error) {

	reqJson, _ := d.Request.ToRequestJson()
	respJson, _ := d.Response.ToResponseJson()

	type mock struct {
		Request  json.RawMessage `json:"request"`
		Response json.RawMessage `json:"response"`
	}

	mockDefinition := mock{
		Request:  reqJson,
		Response: respJson,
	}

	return json.Marshal(mockDefinition)
}
