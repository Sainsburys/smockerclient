package mock

type ResponseBuilder struct {
	response *Response
}

func NewResponseBuilder(httpStatus int) ResponseBuilder {
	response := Response{
		Status: httpStatus,
	}
	return ResponseBuilder{
		response: &response,
	}
}

func (rb ResponseBuilder) Build() Response {
	return *rb.response
}

func (rb ResponseBuilder) AddHeader(key string, values ...string) ResponseBuilder {
	rb.initialiseHeaders()
	rb.response.Headers[key] = values
	return rb
}

func (rb ResponseBuilder) initialiseHeaders() {
	if rb.response.Headers == nil {
		rb.response.Headers = make(map[string][]string, 1)
	}
}

func (rb ResponseBuilder) AddBody(body string) ResponseBuilder {
	rb.response.Body = body
	return rb
}
