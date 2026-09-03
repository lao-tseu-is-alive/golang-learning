package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println(`
which expect one argument : The name of a potential command ! 
Which then returns the path of the files (or links) which would be executed 
in the current environment that are having the given command name.`)
		return
	}
	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			// Is it a regular file?
			if mode.IsRegular() {
				// Is it executable?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					os.Exit(0)
				}
			}
		}
	}
	fmt.Printf("which did not find any command named '%s' in PATH\n", file)
	os.Exit(1)
}
