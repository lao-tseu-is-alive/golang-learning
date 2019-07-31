package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	for i := 1; i <= 6; i++ {
		_, err := os.Create(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
	}
	/*
		To get the filtered file list which corresponds to the given pattern,
		the Glob function from the filepath package can be used.
		For the pattern syntax, see the documentation of the
		filepath.Match function (https://golang.org/pkg/path/filepath/#Match).
	*/

	m, err := filepath.Glob("./test.file[1-3]")
	if err != nil {
		panic(err)
	}

	for _, val := range m {
		fmt.Println(val)
	}

	// Cleanup
	for i := 1; i <= 6; i++ {
		err := os.Remove(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
	}
}
