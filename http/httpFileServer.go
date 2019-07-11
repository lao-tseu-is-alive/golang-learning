package main

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Indent indenting the input by given indent and rune
func IndentByRune(input string, indent int, r rune) string {
	return strings.Repeat(string(r), indent) + input
}

// Indent indenting the input by given indent and rune
func PadByRune(input string, indent int, r rune) string {
	// golog.Info("%q - len= %d, indent= %d",input,len(input), indent)
	if len(input) < (indent - 3) {
		repeat := indent - len(input)
		return input + strings.Repeat(string(r), repeat)
	} else {
		name := input[:(len(input) - 3)]
		return name + string('…') + ".."
	}
}

func getAllFiles(dirPath string) ([]string, error) {
	var files []string
	var dirs []string
	subDirToSkip := ".git"
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			golog.Warn("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			if info.IsDir() && info.Name() == subDirToSkip {
				golog.Info("skipping a dir without errors: %+v \n", info.Name())
				return filepath.SkipDir
			} else {
				dirs = append(dirs, path)
				//fmt.Printf("<DIR>: %q\n", filepath.Base(path))
			}
		} else {
			res := fmt.Sprintf("%s %d", PadByRune(filepath.Base(path), 35, '.'), info.Size())
			files = append(files, res)

		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dirPath, err)
		return nil, err
	}
	if len(files) > 1 {
		sort.Strings(files)
	}
	return files, nil
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	allFiles, err := getAllFiles(initialDirectory)
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // normal header

	if err != nil {
		golog.Err("error getAllFiles(path: %q: %v\n", initialDirectory, err)
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to root !\n")
	for _, file := range allFiles {
		fmt.Fprintf(w, "%s\n<br>", file)
	}
}

var initialDirectory string

func main() {
	// example : https://golang.org/src/net/http/example_test.go
	const defaultHost = "localhost"
	const defaultPort = 8080
	port := defaultPort
	defaultBasePath, err := os.Getwd()
	if err != nil {
		golog.Err("Problem getting current working directory")
		panic(err)
	}
	val, exist := os.LookupEnv("WEB_PORT")
	if !exist {
		flag.IntVar(&port, "port", defaultPort, "port the server will listen to")
	} else {
		golog.Info("Using ENV variable WEB_PORT to listen ")
		port, err = strconv.Atoi(val)
	}
	flag.StringVar(&initialDirectory, "dir", defaultBasePath, "Give the path of the directory you want to serve")
	flag.Parse()
	initialDirectory, err = filepath.Abs(initialDirectory)
	if err != nil {
		golog.Err("Problem getting absolute path of directory: %s", initialDirectory)
		panic(err)
	}
	if _, err := os.Stat(initialDirectory); os.IsNotExist(err) {
		golog.Err("The dir parameter is wrong, %s is not a valid directory", initialDirectory)
		os.Exit(1)
	}

	golog.Info("\nlistening on port : %d\nServing files from: %s", port, initialDirectory)

	http.HandleFunc("/test", filesHandler)
	//http.Handle("/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir(initialDirectory))))
	http.Handle("/", http.FileServer(http.Dir(initialDirectory)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", defaultHost, port), nil))
}
