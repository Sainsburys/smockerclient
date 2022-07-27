package mock

import "github.com/churmd/smockerclient"

type JsonMock struct {
	json string
}

func NewJsonMock(json string) smockerclient.MockDefinition {
	return JsonMock{
		json: json,
	}
}

func (jm JsonMock) MarshalJSON() ([]byte, error) {
	return []byte(jm.json), nil
}
