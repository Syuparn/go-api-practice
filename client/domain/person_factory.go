package domain

type PersonFactory interface {
	Create(name string, age int) (Person, error)
}
