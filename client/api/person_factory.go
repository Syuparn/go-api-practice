package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Syuparn/go-api-practice/client/domain"
)

type PersonFactory struct {
	client *http.Client
}

func NewPersonFactory() *PersonFactory {
	return &PersonFactory{client: &http.Client{}}
}

func (f *PersonFactory) Create(
	name string,
	age int,
) (domain.Person, error) {
	const END_POINT = API_DOMAIN + "/persons"

	age_, err := domain.NewAge(age)
	if err != nil {
		return domain.Person{}, err
	}

	name_, err := domain.NewName(name)
	if err != nil {
		return domain.Person{}, err
	}

	reqJson, err := json.Marshal(newReqPost(name_, age_))
	if err != nil {
		return domain.Person{}, err
	}

	req, err := http.NewRequest(http.MethodPost, END_POINT,
		bytes.NewBuffer(reqJson))
	if err != nil {
		return domain.Person{}, err
	}

	res, err := f.client.Do(req)
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

type _ReqPost struct {
	Age  domain.Age  `json:"age"`
	Name domain.Name `json:"name"`
}

func newReqPost(name domain.Name, age domain.Age) _ReqPost {
	return _ReqPost{Age: age, Name: name}
}

// check if PersonFactory really implements domain.PersonFactory
var (
	_ domain.PersonFactory = NewPersonFactory()
)
