package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Route("/persons", func(r chi.Router) {

		r.Post("/", createPerson)
		r.Get("/", readPersons)
		r.Put("/{personID}", updatePerson)
		r.Delete("/{personID}", deletePerson)
	})

	http.ListenAndServe("localhost:8080", r)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /persons")

	age, name, ok := handlePostReq(w, r)
	if !ok {
		return
	}

	person, err := NewPersonFactory().Create(age, name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errJSON(err)))
		return
	}

	resJSON, err := postResJSON(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errJSON(err)))
		return
	}

	fmt.Printf("send response: %s\n", string(resJSON))
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func readPersons(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /persons")

	persons, err := NewPersonRepository().Read()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errJSON(err)))
		return
	}

	resJSON, err := getResJSON(persons)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errJSON(err)))
		return
	}

	fmt.Printf("send response: %s\n", string(resJSON))
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {}
func deletePerson(w http.ResponseWriter, r *http.Request) {}
