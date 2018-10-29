package main

import (
	"fmt"

	"github.wdf.sap.corp/i332859/learn-go/letsgo/interfaces/person"
	"github.wdf.sap.corp/i332859/learn-go/letsgo/interfaces/utils"
)

func main() {
	p1, err := person.New("a", 0, 2.0)
	if err != nil {
		fmt.Println("Error")
		if err == person.ErrAgeNotValid {
			fmt.Println("Age not valid")
		}
	} else {
		p1.Print()
	}

	p2, err := person.New("b", 1, 2.0)
	if err != nil {
		fmt.Println("Error")
		if err == person.ErrAgeNotValid {
			fmt.Println("Age not valid")
		}
	} else {
		p2.Print()
	}

	p, _ := person.New("c", 1, 2.0)
	printer.DoPrint(p)

	printer.DoPrint(Dummy{})
}

type Dummy struct{}

func (d Dummy) Print() {

}
