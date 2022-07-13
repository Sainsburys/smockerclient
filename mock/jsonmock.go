package mock

type JsonMock struct {
	json string
}

func NewJsonMock(json string) JsonMock {
	return JsonMock{
		json: json,
	}
}

func (jm JsonMock) ToJsonDefinition() []byte {
	return []byte(jm.json)
}

func (jm JsonMock) MarshalJSON() ([]byte, error) {
	return []byte(jm.json), nil
}
