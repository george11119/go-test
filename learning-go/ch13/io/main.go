package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := "hello"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		panic(err)
	}
	fmt.Println(counts)
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
