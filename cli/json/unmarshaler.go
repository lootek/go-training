package json

import "encoding/json"

type JSONUnmarshaler struct{}

func (JSONUnmarshaler) Unmarshal(data []byte) ([]int, error) {
	var v []int
	err := json.Unmarshal(data, &v)

	return v, err
}
