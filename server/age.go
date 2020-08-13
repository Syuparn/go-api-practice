package main

import (
	"fmt"
)

type Age int

func (a Age) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", a)), nil
}

func NewAge(age int) (Age, error) {
	if age < 0 {
		return Age(0), fmt.Errorf("age must be positive")
	}

	return Age(age), nil
}
