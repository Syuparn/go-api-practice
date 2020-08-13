package main

import "encoding/json"

func postResJSON(person Person) ([]byte, error) {
	return json.Marshal(person)
}
