package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
