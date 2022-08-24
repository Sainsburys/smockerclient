package mock

import "encoding/json"

type Request struct {
	Method      string              `json:"method"`
	Path        string              `json:"path"`
	QueryParams map[string][]string `json:"query_params,omitempty"`
	Headers     map[string][]string `json:"headers,omitempty"`
	Body        *RequestBody        `json:"body,omitempty"`
}

type RequestBody struct {
	Matcher string `json:"matcher"`
	Value   string `json:"value"`
}

type Response struct {
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers,omitempty"`
	Body    string              `json:"body,omitempty"`
}

type Definition struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

func NewDefinition(req Request, resp Response) Definition {
	return Definition{
		Request:  req,
		Response: resp,
	}
}

func (d Definition) ToMockDefinitionJson() ([]byte, error) {
	return json.Marshal(d)
}
