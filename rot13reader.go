package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	for i, chr := range string(b) {					
		if unicode.IsLetter(chr) {						
			b[i] = b[i] + 13
			if b[i] > 122 {
				b[i] -= 26
			}
        }					
	}	
	return n, err	
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
