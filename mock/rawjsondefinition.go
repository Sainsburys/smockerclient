package mock

import "github.com/Sainsburys/smockerclient"

type RawJsonDefinition struct {
	json string
}

func NewRawJsonDefinition(json string) smockerclient.MockDefinition {
	return RawJsonDefinition{
		json: json,
	}
}

func (jm RawJsonDefinition) ToMockDefinitionJson() ([]byte, error) {
	return []byte(jm.json), nil
}
