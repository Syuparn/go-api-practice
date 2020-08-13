package view

import (
	"fmt"

	"github.com/Syuparn/go-api-practice/client/domain"
)

const _HEADER_FORMAT = "%36s %12s %3s"
const _RECORD_FORMAT = "%36s %12s %3d"

func ShowPersons(persons []domain.Person) {
	fmt.Printf(_HEADER_FORMAT, "id", "name", "age")
	for _, person := range persons {
		showPerson(person)
	}
}

func showPerson(person domain.Person) {
	fmt.Printf(_RECORD_FORMAT, person.ID.String(), person.Name, person.Age)
}
