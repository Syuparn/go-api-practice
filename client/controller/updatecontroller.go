package controller

import (
	"github.com/Syuparn/go-api-practice/client/domain"
	"github.com/Syuparn/go-api-practice/client/view"
	"github.com/gofrs/uuid"
)

type UpdateController struct {
	personRepository domain.PersonRepository
}

func NewUpdateController(repository domain.PersonRepository) UpdateController {
	return UpdateController{personRepository: repository}
}

func (c *UpdateController) Update(age int, id string, name string) error {
	id_, err := uuid.FromString(id)
	if err != nil {
		return err
	}

	person, err := c.personRepository.Update(id_, name, age)
	if err != nil {
		return err
	}

	view.ShowPersons([]domain.Person{person})
	return nil
}
