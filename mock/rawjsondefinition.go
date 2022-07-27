package mock

import "github.com/churmd/smockerclient"

type RawJsonDefinition struct {
	json string
}

func NewRawJsonDefinition(json string) smockerclient.Mock {
	return RawJsonDefinition{
		json: json,
	}
}

func (jm RawJsonDefinition) ToMockJson() ([]byte, error) {
	return []byte(jm.json), nil
}
