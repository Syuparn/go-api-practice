package api

import (
	"testing"

	"github.com/Syuparn/go-api-practice/client/domain"

	"github.com/gofrs/uuid"
)

var MOCK_UUID0 = uuid.Must(uuid.NewV4())
var MOCK_UUID1 = uuid.Must(uuid.NewV4())
var MOCK_UUID2 = uuid.Must(uuid.NewV4())

func newAge(age int) domain.Age {
	ageObj, err := domain.NewAge(age)
	if err != nil {
		panic(err)
	}
	return ageObj
}

func newName(name string) domain.Name {
	nameObj, err := domain.NewName(name)
	if err != nil {
		panic(err)
	}
	return nameObj
}

func newPerson(name string, age int, id uuid.UUID) domain.Person {
	person, err := domain.NewPerson(name, age, id)
	if err != nil {
		panic(err)
	}
	return person
}

func comparePerson(t *testing.T, actual, expected domain.Person) {
	if actual.Name != expected.Name {
		t.Errorf("wrong name: expected=%s, actual=%s", expected.Name, actual.Name)
	}

	if actual.Age != expected.Age {
		t.Errorf("wrong age: expected=%d, actual=%d", expected.Age, actual.Age)
	}

	if actual.ID != expected.ID {
		t.Errorf("wrong id: expected=%s, actual=%s",
			expected.ID.String(), actual.ID.String())
	}
}
