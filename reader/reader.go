package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	b = b[:cap(b)]
	for i := range b {
		b[i] = byte('A')
	}
	return len(b), nil

}

func main() {
	reader.Validate(MyReader{})
}
