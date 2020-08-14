package domain

import "github.com/gofrs/uuid"

type PersonRepository interface {
	Read() ([]Person, error)
	Update(id uuid.UUID, name string, age int) (Person, error)
}
