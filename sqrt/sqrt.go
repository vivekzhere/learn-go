package main

import (
	"fmt"
)

// Sqrt return the squareroot of x
func Sqrt(x float64) float64 {
	z := float64(1)
	precision := 0.00001
	diff := z*z - x
	for ; diff > precision || diff < -precision; diff = z*z - x {
		z -= diff / (2 * z)
	}
	z -= diff / (2 * z)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
