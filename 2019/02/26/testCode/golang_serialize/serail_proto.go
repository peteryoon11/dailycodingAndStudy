// https://code.tutsplus.com/tutorials/json-serialization-with-golang--cms-30209
package golang_serialize

import (
	"encoding/json"
	"errors"
)

type EnumType int

const (
	Zero EnumType = iota
	One
)

func (e *EnumType) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	value, ok := map[string]EnumType{"Zero": Zero, "One": One}[s]
	if !ok {
		return errors.New("Invalid EnumType value")
	}
	*e = value
	return nil
}

func (e *EnumType) MarshalJSON() ([]byte, error) {
	value, ok := map[EnumType]string{Zero: "Zero", One: "One"}[*e]
	if !ok {
		return nil, errors.New("Invalid EnumType value")
	}
	return json.Marshal(value)
}
