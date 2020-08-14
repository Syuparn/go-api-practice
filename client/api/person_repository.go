package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Syuparn/go-api-practice/client/domain"
	"github.com/gofrs/uuid"
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

func (r *PersonRepository) Update(
	id uuid.UUID,
	name string,
	age int,
) (domain.Person, error) {
	END_POINT := API_DOMAIN + fmt.Sprintf("/persons/%s", id.String())

	age_, err := domain.NewAge(age)
	if err != nil {
		return domain.Person{}, err
	}

	name_, err := domain.NewName(name)
	if err != nil {
		return domain.Person{}, err
	}

	reqJson, err := json.Marshal(newReqPut(name_, age_))
	if err != nil {
		return domain.Person{}, err
	}

	req, err := http.NewRequest(http.MethodPut, END_POINT,
		bytes.NewBuffer(reqJson))
	if err != nil {
		return domain.Person{}, err
	}

	res, err := r.client.Do(req)
	if err != nil {
		return domain.Person{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return domain.Person{}, err
	}

	var person domain.Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		return domain.Person{}, err
	}

	return person, nil
}

type _ReqPut struct {
	Age  domain.Age  `json:"age"`
	Name domain.Name `json:"name"`
}

func newReqPut(name domain.Name, age domain.Age) _ReqPut {
	return _ReqPut{Age: age, Name: name}
}

func (r *PersonRepository) Delete(id uuid.UUID) error {
	END_POINT := API_DOMAIN + fmt.Sprintf("/persons/%s", id.String())

	req, err := http.NewRequest(http.MethodDelete, END_POINT, nil)
	if err != nil {
		return err
	}

	_, err = r.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// check if PersonRepository really implements domain.PersonRepository
var (
	_ domain.PersonRepository = NewPersonRepository()
)
