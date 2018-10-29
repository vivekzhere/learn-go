package person

import (
	"fmt"
)

type Person struct {
	Name   string
	age    int
	salary float64
}

var ErrAgeNotValid error

func New(name string, age int, salary float64) (*Person, error) {
	if age <= 0 {
		return nil, ErrAgeNotValid
	}
	return &Person{
		Name:   name,
		age:    age,
		salary: salary,
	}, nil
}

func init() {
	ErrAgeNotValid = fmt.Errorf("age is not valid")
}

func (p *Person) Print() {
	fmt.Println(p)
}
