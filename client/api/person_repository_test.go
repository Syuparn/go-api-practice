package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Syuparn/go-api-practice/client/domain"
	"gopkg.in/h2non/gock.v1"

	"github.com/gofrs/uuid"
)

var MOCK_UUID0 = uuid.Must(uuid.NewV4())
var MOCK_UUID1 = uuid.Must(uuid.NewV4())
var MOCK_UUID2 = uuid.Must(uuid.NewV4())

func TestPersonRepositoryRead(t *testing.T) {
	tests := []struct {
		response []map[string]string
		expected []domain.Person
	}{
		{
			[]map[string]string{
				map[string]string{
					"name": "Taro",
					"age":  "20",
					"id":   MOCK_UUID0.String(),
				},
			},
			[]domain.Person{
				newPerson("Taro", 20, MOCK_UUID0),
			},
		},
		// multiple response
		{
			[]map[string]string{
				map[string]string{
					"name": "Taro",
					"age":  "20",
					"id":   MOCK_UUID0.String(),
				},
				map[string]string{
					"name": "Hanako",
					"age":  "25",
					"id":   MOCK_UUID1.String(),
				},
			},
			[]domain.Person{
				newPerson("Taro", 20, MOCK_UUID0),
				newPerson("Hanako", 25, MOCK_UUID1),
			},
		},
	}

	for i, tt := range tests {
		// move range scope to local scope
		tt := tt
		i := i

		t.Run(fmt.Sprintf("%d: response: %s", i, tt.response),
			func(t *testing.T) {
				testPersonRepositoryRead(t, tt.response, tt.expected)
			})
	}
}

func TestPersonRepositoryUpdate(t *testing.T) {
	tests := []struct {
		age      int
		id       uuid.UUID
		name     string
		expected domain.Person
	}{
		{
			20,
			MOCK_UUID0,
			"Taro",
			newPerson("Taro", 20, MOCK_UUID0),
		},
	}

	for i, tt := range tests {
		// move range scope to local scope
		tt := tt
		i := i

		t.Run(fmt.Sprintf("%d: id %s, age %d, name %s", i, tt.id, tt.age, tt.name),
			func(t *testing.T) {
				testPersonRepositoryUpdate(
					t, tt.id, tt.name, tt.age, tt.expected)
			})
	}
}

func TestPersonRepositoryDelete(t *testing.T) {
	testPersonRepositoryDelete(t, MOCK_UUID0)
}

func testPersonRepositoryRead(
	t *testing.T,
	response []map[string]string,
	expected []domain.Person,
) {
	defer gock.Off()

	type Response struct {
		Persons []map[string]string `json:"persons"`
	}

	client := &http.Client{}
	gock.New(API_DOMAIN).
		Get("/persons").
		Reply(200).
		JSON(Response{Persons: response})
	gock.InterceptClient(client)

	personRepository := &PersonRepository{client: client}
	persons, err := personRepository.Read()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(persons) != len(expected) {
		t.Errorf("wrong response length: expected=%d(%v), got=%d(%v)",
			len(expected), expected, len(persons), persons)
		return
	}

	for i, person := range persons {
		comparePerson(t, person, expected[i])
	}
}

func testPersonRepositoryUpdate(
	t *testing.T,
	id uuid.UUID,
	name string,
	age int,
	expected domain.Person,
) {
	defer gock.Off()

	client := &http.Client{}
	gock.New(API_DOMAIN).
		Put(fmt.Sprintf("/persons/%s", id.String())).
		Reply(200).
		JSON(map[string]string{
			"id":   MOCK_UUID0.String(),
			"name": string(name),
			"age":  fmt.Sprintf("%d", age),
		})
	gock.InterceptClient(client)

	personRepository := &PersonRepository{client: client}
	person, err := personRepository.Update(id, name, age)

	if err != nil {
		t.Fatalf(err.Error())
	}

	comparePerson(t, person, expected)
}

func testPersonRepositoryDelete(t *testing.T, id uuid.UUID) {
	defer gock.Off()

	client := &http.Client{}
	gock.New(API_DOMAIN).
		Delete(fmt.Sprintf("/persons/%s", id.String())).
		Reply(204)
	gock.InterceptClient(client)

	personRepository := &PersonRepository{client: client}
	err := personRepository.Delete(id)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

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
