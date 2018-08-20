package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := range b {
		switch {
		case b[i] >= 'a' && b[i] <= 'z':
			b[i] = (b[i]-'a'+13)%26 + 'a'
		case b[i] >= 'A' && b[i] <= 'Z':
			b[i] = (b[i]-'A'+13)%26 + 'A'
		default:
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	s2 := strings.NewReader("You cracked the code!")
	r2 := rot13Reader{s2}
	io.Copy(os.Stdout, &r2)
}
