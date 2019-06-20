package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"log"
	"os"
	"path"
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

func main() {
	dataFile := path.Join(defaultBasePath, defaultTextFile)
	allText, err := GetTextInFile(dataFile)
	if err != nil {
		golog.Err(allText)
		log.Fatal(allText)
	}
	allLines := strings.Split(allText, "\n")
	numLines := len(allLines)
	//fmt.Printf("%+v\n\n",allText)
	fmt.Printf("%+v\n\n", allLines)
	fmt.Printf("#File: %s contains %d lines\n", defaultTextFile, numLines)
	fmt.Println(allLines[0])

}
