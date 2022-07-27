package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Request struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	QueryParams MultiMap `json:"query_params,omitempty"`
	Headers     MultiMap `json:"headers,omitempty"`
	Body        *Body    `json:"body,omitempty"`
}

func NewRequest(method, path string) Request {
	return Request{
		Method: method,
		Path:   path,
	}
}

func (r *Request) AddQueryParam(key, value string) {
	if r.QueryParams == nil {
		r.QueryParams = MultiMap{}
	}

	r.QueryParams[key] = value
}

func (r *Request) AddHeader(key, value string) {
	if r.Headers == nil {
		r.Headers = MultiMap{}
	}

	r.Headers[key] = value
}

func (r *Request) AddJsonBody(jsonBody string) error {
	compactJsonBody, err := compactJson(jsonBody)
	if err != nil {
		return fmt.Errorf("unable to compact body json. %w", err)
	}

	body := Body{
		Matcher: "ShouldEqualJSON",
		Value:   compactJsonBody,
	}
	r.Body = &body

	return nil
}

func compactJson(jsonObject string) (string, error) {
	compactJsonObject := new(bytes.Buffer)
	err := json.Compact(compactJsonObject, []byte(jsonObject))
	if err != nil {
		return "", err
	}
	return compactJsonObject.String(), nil
}

type MultiMap map[string]string

func (qp MultiMap) MarshalJSON() ([]byte, error) {
	paramsAsJson := qp.combineKeyValuePairsForJson()
	multiMapJson := "{" + paramsAsJson + "}"
	return []byte(multiMapJson), nil
}

func (qp MultiMap) combineKeyValuePairsForJson() string {
	params := make([]string, 0)

	for key, value := range qp {
		pair := `"` + key + `":"` + value + `"`
		params = append(params, pair)
	}

	return strings.Join(params, ",")
}

type Body struct {
	Matcher string `json:"matcher"`
	Value   string `json:"value"`
}
