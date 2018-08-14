package main

import (
	"golang.org/x/tour/pic"
)

// Pic returns a slice of length dy, each element is a slice of dx numbers
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
		for j := range result[i] {
			result[i][j] = uint8((i * j) ^ (i - j))
		}
	}
	return result
}

func main() {
	pic.Show(Pic)
}
