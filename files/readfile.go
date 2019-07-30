package main

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		panic("Expecting at least one argument indicating the file you want to read !")
	}
	var file2read = ""
	const fileHelp = "filename you want to read (and print the content)"
	if len(args) > 2 {
		flag.StringVar(&file2read, "file", "", fileHelp)
		flag.StringVar(&file2read, "f", "", fileHelp+"(shorthand)")
		flag.Parse()
	} else {
		file2read = args[1]
	}
	if len(file2read) > 0 {
		f, err := os.Open(file2read)
		if err != nil {
			golog.Err("Error trying to open file : %s \n Error : %v\n", file2read, err)
			log.Fatal(err)
		}
		defer f.Close()

		c, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		fmt.Printf("### File content [type : %T]###\n%s\n", c, string(c))

	} else {
		panic("Invalid File Name !")
	}

}
