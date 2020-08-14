package main

import "encoding/json"

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
