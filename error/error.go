package main

import (
	"fmt"
	"reflect"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt return the squareroot of x
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	precision := 0.00001
	diff := z*z - x
	for ; diff > precision || diff < -precision; diff = z*z - x {
		z -= diff / (2 * z)
	}
	z -= diff / (2 * z)
	return z, nil
}

func main() {
	_, err := Sqrt(-2)
	//fmt.Println(err.(type))
	fmt.Println(reflect.TypeOf(err))
	switch err.(type) {
	case ErrNegativeSqrt:
		fmt.Println("custom type")
	case error:
		fmt.Println("general type")
	}
	_, err = Sqrt(2)
	fmt.Println(reflect.TypeOf(err))
	switch err.(type) {
	case ErrNegativeSqrt:
		fmt.Println("custom type")
	case error:
		fmt.Println("general type")
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
