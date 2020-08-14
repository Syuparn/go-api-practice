package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
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

func updatePerson(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "personID")
	fmt.Println("PUT /persons/" + idStr)

	id, err := uuid.FromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errJSON(err)))
		return
	}

	age, name, ok := handlePutReq(w, r)
	if !ok {
		return
	}

	repo := NewPersonRepository()
	ok = repo.Exists(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errJSON(fmt.Errorf("id not found"))))
		return
	}

	person, err := repo.Update(id, age, name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errJSON(err)))
		return
	}

	resJSON, err := putResJSON(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errJSON(err)))
		return
	}

	fmt.Printf("send response: %s\n", string(resJSON))
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "personID")
	fmt.Println("DELETE /persons/" + idStr)

	id, err := uuid.FromString(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errJSON(err)))
		return
	}

	err = NewPersonRepository().Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errJSON(err)))
		return
	}

	fmt.Println("send response")
	w.WriteHeader(http.StatusNoContent)
}
