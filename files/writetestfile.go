package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const defaultFileName = "test.txt"

	f, err := os.OpenFile(defaultFileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	n, err := io.WriteString(f, "Test string")
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes in file : %s\n", n, defaultFileName)

}
