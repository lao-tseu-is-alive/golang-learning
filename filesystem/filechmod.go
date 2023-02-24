package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Obtain current permissions // -rw-r--r--
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File permissions %v\n", fi.Mode())

	/*
		The shortest way to change the permissions is by using the os.Chmod function,
		which does the same, but you do not need to obtain the File type in the code.
	*/

	// Change permissions	// -rwxrwxrwx
	err = f.Chmod(0777)
	if err != nil {
		panic(err)
	}
	fi, err = f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File permissions %v\n", fi.Mode())

}
