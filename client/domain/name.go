package domain

import (
	"fmt"
)

type Name string

func (n Name) MarshalJSON() ([]byte, error) {
	// NOTE: value in json must be double-quoted!
	return []byte(fmt.Sprintf(`"%s"`, n)), nil
}

func NewName(name string) (Name, error) {
	if len(name) == 0 {
		return Name(""), fmt.Errorf("name must not be empty")
	}

	return Name(name), nil
}
