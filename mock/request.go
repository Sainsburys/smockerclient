package mock

import (
	"strings"
)


type Request struct {
	Method string `json:"method"`
	Path string `json:"path"`
	QueryParams QueryParams `json:"query_params,omitempty"`
}

func NewRequest(method, path string) Request {
	return Request{
		Method: method,
		Path: path,
	}
}

func (r *Request) AddQueryParam(key, value string) {
	if r.QueryParams == nil {
		r.QueryParams = QueryParams{}
	}

	r.QueryParams[key] = value
}

type QueryParams map[string]string

func (qp QueryParams) MarshalJSON() ([]byte, error) {
	paramsAsJson := qp.combineKeyValuePairsForJson()
	json := "{" + paramsAsJson + "}"
	return []byte(json), nil
}

func (qp QueryParams) combineKeyValuePairsForJson() string {
	params := make([]string, 0)

	for key, value := range qp {
		pair := `"` + key + `":"` + value + `"`
		params = append(params, pair)
	}

	return strings.Join(params, ",")
}