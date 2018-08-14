package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount returns the count of all words in the string
func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	wc := make(map[string]int)
	for _, word := range fields {
		wc[word]++
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
