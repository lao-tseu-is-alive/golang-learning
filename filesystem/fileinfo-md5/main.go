package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"io"
	"log"
	"os"
)

func main() {

	args := os.Args
	if len(args) == 1 {
		panic("Expecting at least one argument indicating the file you want to get information !")
	}
	var file2read = ""
	const fileHelp = "filename you want to get information"
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

		fi, err := f.Stat()
		if err != nil {
			golog.Err("Error trying to get info (Stat) of  file : %s \n Error : %v\n", file2read, err)
			panic(err)
		}
		fmt.Printf("### File information for :###\n%s\n", file2read)
		fmt.Printf("File name   : %v\n", fi.Name())
		fmt.Printf("Is Directory: %t\n", fi.IsDir())
		fmt.Printf("Size  : %d\n", fi.Size())
		fmt.Printf("Mode  : %v\n", fi.Mode())
		fmt.Printf("Time  : %v\n", fi.ModTime())

		// now will get the md5sum of the file
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			golog.Err("Error trying to get md5 ! Error : %v\n", err)
			log.Fatal(err)
		}

		fmt.Printf("md5sum: %x\n", h.Sum(nil))

	} else {
		panic("Invalid File Name !")
	}

}
