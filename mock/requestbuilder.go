package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type RequestBuilder struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	QueryParams MultiMap `json:"query_params,omitempty"`
	Headers     MultiMap `json:"headers,omitempty"`
	Body        *Body    `json:"body,omitempty"`
}

func NewRequestBuilder(method, path string) RequestBuilder {
	return RequestBuilder{
		Method: method,
		Path:   path,
	}
}

func (r *RequestBuilder) AddQueryParam(key string, values ...string) {
	if r.QueryParams == nil {
		r.QueryParams = MultiMap{}
	}

	r.QueryParams[key] = values
}

func (r *RequestBuilder) AddHeader(key string, values ...string) {
	if r.Headers == nil {
		r.Headers = MultiMap{}
	}

	r.Headers[key] = values
}

func (r *RequestBuilder) AddJsonBody(jsonBody string) error {
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

type MultiMap map[string][]string

func (mm MultiMap) MarshalJSON() ([]byte, error) {
	paramsAsJson, err := mm.combineKeyValuePairsForJson()
	if err != nil {
		return nil, fmt.Errorf("unable json convert mutlimap %+v. %w", mm, err)
	}

	multiMapJson := "{" + paramsAsJson + "}"
	return []byte(multiMapJson), nil
}

func (mm MultiMap) combineKeyValuePairsForJson() (string, error) {
	params := make([]string, 0)

	for key, values := range mm {
		jsonValues, err := json.Marshal(values)
		if err != nil {
			return "", err
		}

		pair := `"` + key + `":` + string(jsonValues)
		params = append(params, pair)
	}

	return strings.Join(params, ","), nil
}

type Body struct {
	Matcher string `json:"matcher"`
	Value   string `json:"value"`
}
