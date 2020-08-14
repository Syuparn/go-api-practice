package domain

type PersonRepository interface {
	Read() ([]Person, error)
}
