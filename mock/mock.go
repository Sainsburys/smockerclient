package mock

import (
	"encoding/json"
	"fmt"
)

type Request json.Marshaler
type Response json.Marshaler

type Definition struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

func (m Definition) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
