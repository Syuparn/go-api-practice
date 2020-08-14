package controller

import (
	"github.com/Syuparn/go-api-practice/client/domain"
	"github.com/Syuparn/go-api-practice/client/view"
)

type ReadController struct {
	personRepository domain.PersonRepository
}

func NewReadController(repository domain.PersonRepository) ReadController {
	return ReadController{personRepository: repository}
}

func (c *ReadController) Read() error {
	persons, err := c.personRepository.Read()
	if err != nil {
		return err
	}

	view.ShowPersons(persons)
	return nil
}
