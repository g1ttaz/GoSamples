package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case b >= 'A' && b <= 'M':
		b += 13
	case b >= 'a' && b <= 'm':
		b += 13
	case b >= 'N' && b <= 'Z':
		b -= 13
	case b >= 'n' && b <= 'z':
		b -= 13
	}
	return b
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
	if err != nil {
		return n, err
	}
	
	for i := range(b) {
		b[i] = rot13(b[i])
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

