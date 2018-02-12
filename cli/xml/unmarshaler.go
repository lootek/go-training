package xml

import "encoding/xml"

type XMLUnmarshaler struct{}

func (XMLUnmarshaler) Unmarshal(data []byte) ([]int, error) {
	var v []int
	err := xml.Unmarshal(data, &v)

	return v, err
}
