package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type RequestBody struct {
	Matcher string
	Value   string
}

func NewJsonRequestBody(jsonBody string) RequestBody {
	return RequestBody{
		Matcher: "ShouldEqualJSON",
		Value:   jsonBody,
	}
}

func (b RequestBody) MarshalJSON() ([]byte, error) {
	compactValue, err := compactJson(b.Value)
	if err != nil {
		return nil, fmt.Errorf("unable to compact body json %s. %w", b.Value, err)
	}

	type compactBody struct {
		Matcher string `json:"matcher"`
		Value   string `json:"value"`
	}

	cb := compactBody{
		Matcher: b.Matcher,
		Value:   compactValue,
	}

	return json.Marshal(cb)
}

func compactJson(jsonObject string) (string, error) {
	compactJsonObject := new(bytes.Buffer)
	err := json.Compact(compactJsonObject, []byte(jsonObject))
	if err != nil {
		return "", err
	}
	return compactJsonObject.String(), nil
}
