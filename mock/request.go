package mock

import (
	"strings"
)

type Request struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	QueryParams MultiMap `json:"query_params,omitempty"`
	Headers     MultiMap `json:"headers,omitempty"`
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

type MultiMap map[string]string

func (qp MultiMap) MarshalJSON() ([]byte, error) {
	paramsAsJson := qp.combineKeyValuePairsForJson()
	json := "{" + paramsAsJson + "}"
	return []byte(json), nil
}

func (qp MultiMap) combineKeyValuePairsForJson() string {
	params := make([]string, 0)

	for key, value := range qp {
		pair := `"` + key + `":"` + value + `"`
		params = append(params, pair)
	}

	return strings.Join(params, ",")
}
