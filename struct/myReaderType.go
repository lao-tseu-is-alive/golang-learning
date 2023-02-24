package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"io"
)

type MyReader struct{}

var Count int = 0

const MaxA = 9

// return a maximum of MaxA 'A'
func (MyReader) Read(b []byte) (int, error) {
	i := 0
	for i = range b {
		golog.Info("in Reader at index %d", i)
		Count += 1
		if Count > MaxA {
			b[i] = 'X'
			return i, io.EOF
		} else {
			b[i] = 'A'
		}
	}
	return i, nil
}

func main() {
	r := MyReader{}
	b := make([]byte, 40)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

}
