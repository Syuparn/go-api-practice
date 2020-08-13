package domain

type PersonFactory interface {
	Create(name Name, age Age) (Person, error)
}
