package mock

import (
	"encoding/json"
	"net/http"
)

type ResponseBuilder struct {
	Status int `json:"status"`
}

func NewResponseBuilder() ResponseBuilder {
	return ResponseBuilder{
		Status: http.StatusOK,
	}
}

func NewNoContentResponse() ResponseBuilder {
	return ResponseBuilder{
		Status: http.StatusNoContent,
	}
}

func (rb ResponseBuilder) ToResponseJson() ([]byte, error) {
	return json.Marshal(rb)
}
