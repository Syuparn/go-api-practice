package main

import (
	"encoding/json"
	"fmt"
)

func postResJSON(person Person) ([]byte, error) {
	return json.Marshal(person)
}

func getResJSON(persons []Person) ([]byte, error) {
	type Response struct {
		Persons []Person `json:"persons"`
	}
	res := Response{Persons: persons}
	return json.Marshal(res)
}

func putResJSON(person Person) ([]byte, error) {
	return json.Marshal(person)
}

func errJSON(err error) string {
	resJSON := fmt.Sprintf(`{"error": "%s"}`, err.Error())
	fmt.Printf("send response: %s\n", string(resJSON))
	return resJSON
}
