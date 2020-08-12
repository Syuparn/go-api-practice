package domain

import (
	"encoding/json"

	"github.com/gofrs/uuid"
)

type Person struct {
	Age  Age       `json:"age"`
	ID   uuid.UUID `json:"id"`
	Name Name      `json:"name"`
}

func (p *Person) UnmarshalJSON(data []byte) error {
	// NOTE: json.Numberなら{age: 20}, {age: "20"}どちらでもパースできる
	// (gockのテストはstringがレスポンスになっているため対応)
	type Response struct {
		Age  json.Number `json:"age"`
		ID   uuid.UUID   `json:"id"`
		Name string      `json:"name"`
	}

	var res Response
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}

	newName, err := NewName(res.Name)
	if err != nil {
		return err
	}
	p.Name = newName

	ageInt, err := res.Age.Int64()
	if err != nil {
		return err
	}

	newAge, err := NewAge(int(ageInt))
	if err != nil {
		return err
	}
	p.Age = newAge

	p.ID = res.ID

	return nil
}

func NewPerson(name string, age int, id uuid.UUID) (Person, error) {
	nameObj, err := NewName(name)
	if err != nil {
		return Person{}, err
	}

	ageObj, err := NewAge(age)
	if err != nil {
		return Person{}, err
	}

	return Person{Age: ageObj, ID: id, Name: nameObj}, nil
}
