package main

import "github.com/gofrs/uuid"

// FIXME: make persistent!
var personDB = []Person{}

type PersonFactory struct{}

func NewPersonFactory() *PersonFactory {
	return &PersonFactory{}
}

func (f *PersonFactory) Create(age int, name string) (Person, error) {
	// NOTE: panic very rarely
	id := uuid.Must(uuid.NewV4())
	person, err := NewPerson(name, age, id)
	if err != nil {
		return Person{}, err
	}

	// register to db
	// NOTE: currently no persistence
	personDB = append(personDB, person)

	return person, nil
}
