package main

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	defaultBasePath, err := os.Getwd()
	if err != nil {
		golog.Err("Problem getting current working directory")
		panic(err)
	}

	directory := flag.String("dir", defaultBasePath, "Give the path of the directory you want to serve")
	flag.Parse()
	*directory, err = filepath.Abs(*directory)
	if err != nil {
		golog.Err("Problem getting absolute path of directory: %s", *directory)
		panic(err)
	}
	if _, err := os.Stat(*directory); os.IsNotExist(err) {
		golog.Err("The dir parameter is wrong, %s is not a valid directory", *directory)
		os.Exit(1)
	}

	golog.Info("\nlisting files from: %s", *directory)

	fmt.Println("List by ReadDir")
	listDirByReadDir(*directory)
	fmt.Println()
	fmt.Println("List by Walk")
	listDirByWalk(*directory)
}

func listDirByWalk(path string) {
	err := filepath.Walk(path, func(wPath string, info os.FileInfo,
		err error) error {

		// Walk the given dir
		// without printing out.
		if wPath == path {
			return nil
		}

		/*
			subDirToSkip := ".git"
			if info.IsDir() && info.Name() == subDirToSkip {
				golog.Info("skipping a dir without errors: %+v \n", info.Name())
				return filepath.SkipDir
			}
		*/

		// If given path is folder
		// stop list recursively and print as folder.
		if info.IsDir() {
			fmt.Printf("DIR [%s]\n", wPath)
			return filepath.SkipDir
		}

		// Print file name
		if wPath != path {
			fmt.Println(filepath.Base(wPath)) // only base filename
			//fmt.Println(wPath) // full path
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %s: %v\n", path, err)
	}
}

func listDirByReadDir(path string) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("DIR [%s]\n", val.Name())
		} else {
			fmt.Println(val.Name())
		}
	}
}
