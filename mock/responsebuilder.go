package mock

type ResponseBuilder struct {
	response Response
}

func NewResponseBuilder(httpStatus int) ResponseBuilder {
	return ResponseBuilder{
		response: Response{
			Status: httpStatus,
		},
	}
}

func (rb ResponseBuilder) Build() Response {
	return rb.response
}

func (rb *ResponseBuilder) AddHeader(key string, values ...string) {
	rb.initialiseHeaders()
	rb.response.Headers[key] = values
}

func (rb *ResponseBuilder) initialiseHeaders() {
	if rb.response.Headers == nil {
		rb.response.Headers = make(map[string][]string, 1)
	}
}

func (rb *ResponseBuilder) AddBody(body string) {
	rb.response.Body = body
}
