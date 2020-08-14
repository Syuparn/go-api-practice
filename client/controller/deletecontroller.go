package controller

import (
	"github.com/Syuparn/go-api-practice/client/domain"
	"github.com/Syuparn/go-api-practice/client/view"
	"github.com/gofrs/uuid"
)

type DeleteController struct {
	personRepository domain.PersonRepository
}

func NewDeleteController(repository domain.PersonRepository) DeleteController {
	return DeleteController{personRepository: repository}
}

func (c *DeleteController) Delete(id string) error {
	id_, err := uuid.FromString(id)
	if err != nil {
		return err
	}

	err = c.personRepository.Delete(id_)
	if err != nil {
		return err
	}

	view.ShowSuccessDelete()
	return nil
}
