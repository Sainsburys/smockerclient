package mock

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MultiMap map[string][]string

func (mm MultiMap) MarshalJSON() ([]byte, error) {
	paramsAsJson, err := mm.combineKeyValuePairsForJson()
	if err != nil {
		return nil, fmt.Errorf("unable json convert mutlimap %+v. %w", mm, err)
	}

	multiMapJson := "{" + paramsAsJson + "}"
	return []byte(multiMapJson), nil
}

func (mm MultiMap) combineKeyValuePairsForJson() (string, error) {
	params := make([]string, 0)

	for key, values := range mm {
		jsonValues, err := json.Marshal(values)
		if err != nil {
			return "", err
		}

		pair := `"` + key + `":` + string(jsonValues)
		params = append(params, pair)
	}

	return strings.Join(params, ","), nil
}
