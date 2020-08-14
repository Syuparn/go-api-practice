package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Syuparn/go-api-practice/client/domain"
)

const API_DOMAIN = "http://localhost:8080"

type PersonRepository struct {
	client *http.Client
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{client: &http.Client{}}
}

func (r *PersonRepository) Read() ([]domain.Person, error) {
	const END_POINT = API_DOMAIN + "/persons"

	req, err := http.NewRequest(http.MethodGet, END_POINT, nil)
	if err != nil {
		return []domain.Person{}, err
	}

	res, err := r.client.Do(req)
	if err != nil {
		return []domain.Person{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []domain.Person{}, err
	}

	type Response struct {
		Persons []domain.Person `json:"persons"`
	}
	var resStruct Response

	err = json.Unmarshal(body, &resStruct)
	if err != nil {
		return []domain.Person{}, err
	}

	return resStruct.Persons, nil
}

// check if PersonRepository really implements domain.PersonRepository
var (
	_ domain.PersonRepository = NewPersonRepository()
)
