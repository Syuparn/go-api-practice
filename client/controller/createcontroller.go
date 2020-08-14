package controller

import (
	"github.com/Syuparn/go-api-practice/client/domain"
	"github.com/Syuparn/go-api-practice/client/view"
)

type CreateController struct {
	personFactory domain.PersonFactory
}

func NewCreateController(factory domain.PersonFactory) CreateController {
	return CreateController{personFactory: factory}
}

func (c *CreateController) Create(age int, name string) error {
	person, err := c.personFactory.Create(name, age)
	if err != nil {
		return err
	}

	view.ShowPersons([]domain.Person{person})
	return nil
}
