package main

import "github.com/gofrs/uuid"

type PersonRepository struct{}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{}
}

func (f *PersonRepository) Read() ([]Person, error) {
	persons := []Person{}

	// NOTE: currently no persistence
	for _, person := range personDB {
		persons = append(persons, person)
	}

	return persons, nil
}

func (f *PersonRepository) Update(id uuid.UUID, age int, name string) (Person, error) {
	person, err := NewPerson(name, age, id)
	if err != nil {
		return Person{}, err
	}

	personDB[id] = person
	return person, nil
}

func (f *PersonRepository) Exists(id uuid.UUID) bool {
	_, ok := personDB[id]
	return ok
}
