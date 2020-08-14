package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func handlePostReq(w http.ResponseWriter, r *http.Request) (int, string, bool) {
	type Req struct {
		Age  json.Number `json:"age"`
		Name string      `json:"name"`
	}

	body, err := extractBody(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 0, "", false
	}

	var req Req
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errJSON(err)))
		return 0, "", false
	}

	age, err := req.Age.Int64()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errJSON(err)))
		return 0, "", false
	}

	return int(age), req.Name, true
}

func handlePutReq(w http.ResponseWriter, r *http.Request) (int, string, bool) {
	// exactly same as POST request
	return handlePostReq(w, r)
}

func extractBody(r *http.Request) ([]byte, error) {
	fmt.Print("request body: ")

	bodyLen, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return []byte{}, err
	}

	body := make([]byte, bodyLen)
	_, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		return []byte{}, err
	}

	fmt.Println(string(body))
	return body, nil
}
