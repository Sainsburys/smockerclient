package mock

type RequestBody struct {
	Matcher string `json:"matcher"`
	Value   string `json:"value"`
}

func NewJsonRequestBody(jsonBody string) RequestBody {
	return RequestBody{
		Matcher: "ShouldEqualJSON",
		Value:   jsonBody,
	}
}
