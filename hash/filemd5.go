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
		panic("Expecting at least one argument indicating the file you want to get Md5 hash !")
	}
	var file2hash = ""
	const fileHelp = "filename you want to get the md5 hash"
	const verboseHelp = "allow to get more verbose feedback."
	var isVerbose = false
	if len(args) > 2 {
		flag.StringVar(&file2hash, "file", "", fileHelp)
		flag.StringVar(&file2hash, "f", "", fileHelp+"(shorthand)")
		flag.BoolVar(&isVerbose, "verbose", false, verboseHelp)
		flag.BoolVar(&isVerbose, "v", false, verboseHelp+"(shorthand)")
		flag.Parse()
	} else {
		file2hash = args[1]
	}
	if len(file2hash) > 0 {
		if isVerbose {
			golog.Info("Will try to find md5 hash of file : %s", file2hash)
		}
		f, err := os.Open(file2hash)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%x  %s\n", h.Sum(nil), file2hash)

	} else {
		panic("Invalid File Name !")
	}
}
