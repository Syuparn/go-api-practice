package main

type PersonRepository struct{}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{}
}

func (f *PersonRepository) Read() ([]Person, error) {
	persons := []Person{}

	// NOTE: currently no persistence
	for _, person := range personDB {
		persons = append(persons, person)
	}

	return persons, nil
}
