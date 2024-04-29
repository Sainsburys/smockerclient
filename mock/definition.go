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

type Context struct {
	Times int `json:"times,omitempty"`
}

type Definition struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
	Context  *Context `json:"context,omitempty"`
}

func NewDefinition(req Request, resp Response, contextOptions ...ContextOption) Definition {
	def := Definition{
		Request:  req,
		Response: resp,
	}

	var context *Context
	for _, fn := range contextOptions {
		context = fn(context)
	}

	if context != nil {
		def.Context = context
	}

	return def
}

func (d Definition) ToMockDefinitionJson() ([]byte, error) {
	return json.Marshal(d)
}
