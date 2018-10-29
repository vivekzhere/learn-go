package printer

type Printable interface {
	Print()
}

func DoPrint(p Printable) {
	p.Print()
}
