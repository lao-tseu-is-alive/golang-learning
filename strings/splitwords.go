package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

const defaultTextFile = "VoltaireCandide.txt"
const defaultBasePath = "data"

func GetTextInFile(filename string) (string, error) {
	// variants of reading : file https://kgrz.io/reading-files-in-go-an-overview.html
	textError := "## ERROR "
	file, err := os.Open(filename)
	if err != nil {
		textError += "OPENING FILE"
		fmt.Println(textError, err)
		return textError, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		textError += "GETTING INFO ON FILE"
		fmt.Println(textError, err)
		return textError, err
	}
	fileSize := fileinfo.Size()
	// fmt.Printf("The FILE is %v bytes\n", fileSize)
	buffer := make([]byte, fileSize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		textError += "READING FILE "
		fmt.Printf(" ERROR READING after %v bytes\n", bytesread)
		fmt.Println(textError, err)
		return textError, err
	}
	// fmt.Printf(" SUCCESSFULLY READ %v \n", bytesread)
	return string(buffer), nil
}

// readLines reads a whole file into memory and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func WordCount(text string) map[string]int {
	reWords := regexp.MustCompile("\\p{L}+")
	words := reWords.FindAllString(text, -1)
	counts := make(map[string]int, len(words))
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func main() {
	dataFile := flag.String("file", path.Join(defaultBasePath, defaultTextFile), "Give the filename you want to get the word list")
	minWordLength := flag.Int("min-length", 1, "minimum lenght of words to keep (default is 1)")
	flag.Parse()
	allLines, err := readLines(*dataFile)
	if err != nil {
		golog.Err("doing readLines %s", err)
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range allLines {
		if len(strings.TrimSpace(line)) > 0 {
			// https://www.regular-expressions.info/unicode.html
			reWords := regexp.MustCompile("(\\p{L}\\p{M}*)+")
			words := reWords.FindAllString(line, -1)
			onlyBiggerWords := []string{}
			for _, word := range words {
				if len(word) > *minWordLength {
					onlyBiggerWords = append(onlyBiggerWords, word)
				}
			}
			wList := strings.Join(onlyBiggerWords, ", ")
			fmt.Printf("%d : %s <|||>[%s]\n", i, line, wList)
		}
	}

	numLines := len(allLines)
	// print all file :
	fmt.Printf("## File: %s contains %d lines\n", *dataFile, numLines)

}
