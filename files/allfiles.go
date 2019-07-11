package files

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Indent indenting the input by given indent and rune
func IndentByRune(input string, indent int, r rune) string {
	return strings.Repeat(string(r), indent) + input
}

// Indent indenting the input by given indent and rune
func PadByRune(input string, indent int, r rune) string {
	//golog.Info("%q - len= %d, indent= %d", input, len(input), indent)
	if len(input) < (indent - 3) {
		repeat := indent - len(input)
		return input + strings.Repeat(string(r), repeat)
	} else {
		name := input[:(len(input) - 3)]
		return name + string('…') + ".."
	}
}

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

	var files []string
	//var dirs []string
	//subDirToSkip := ".git"
	err = filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			golog.Warn("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		/*
			if info.IsDir() && info.Name() == subDirToSkip {
				golog.Info("skipping a dir without errors: %+v \n", info.Name())
				return filepath.SkipDir
			}
		*/
		if info.IsDir() {
			// dirs = append(dirs, path)
			//fmt.Printf("<DIR>: %q\n", filepath.Base(path))
		} else {
			//res := fmt.Sprintf("%s %d", PadByRune(filepath.Base(path), 35, '.'), info.Size())
			//res := filepath.Base(path)
			files = append(files, path)

		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", *directory, err)
		return
	}
	/*
		if len(dirs) > 1 {
			sort.Strings(dirs)
		}
		for _, dir := range dirs {
			fmt.Println(dir)
		}
	*/
	if len(files) > 1 {
		sort.Strings(files)
	}
	for _, file := range files {
		fmt.Println(file)
	}

}
