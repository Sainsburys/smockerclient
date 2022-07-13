package mock

import "fmt"

type Mock struct{}

func (m Mock) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
