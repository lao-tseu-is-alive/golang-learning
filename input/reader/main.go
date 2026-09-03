package main

import (
	"fmt"
	"os"
)

// Try: echo 'hello jimmy!' | go run ./input/reader

func main() {

	for {
		data := make([]byte, 8)
		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
			process(data)
		} else {
			break
		}
	}

}

func process(data []byte) {
	fmt.Printf("Received: %X %s\n", data, string(data))
}
