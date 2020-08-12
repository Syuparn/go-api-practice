package api

import (
	"domain"
	"fmt"
	"net/http"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestPersonFactoryCreate(t *testing.T) {
	const SAMPLE_ID = 1

	tests := []struct {
		age      domain.Age
		name     domain.Name
		expected domain.Person
	}{
		{
			name:     newName("Taro"),
			age:      newAge(20),
			expected: newPerson("Taro", 20, MOCK_UUID0),
		},
	}

	for i, tt := range tests {
		// move range scope to local scope
		tt := tt
		i := i

		t.Run(fmt.Sprintf("%d: age %d, name %s", i, tt.age, tt.name),
			func(t *testing.T) {
				testPersonFactoryCreate(t, tt.age, tt.name, tt.expected)
			})
	}
}

func testPersonFactoryCreate(
	t *testing.T,
	age domain.Age,
	name domain.Name,
	expected domain.Person,
) {
	defer gock.Off()

	client := &http.Client{}
	gock.New(API_DOMAIN).
		Post("/persons").
		Reply(200).
		JSON(map[string]string{
			"name": string(name),
			"age":  fmt.Sprintf("%d", age),
			"id":   MOCK_UUID0.String(),
		})
	gock.InterceptClient(client)

	personFactory := &PersonFactory{client: client}
	person, err := personFactory.Create(name, age)

	if err != nil {
		t.Fatalf(err.Error())
	}

	comparePerson(t, person, expected)
}
