package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/goutils"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	const filePath = "win1252.txt"

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read all in raw form.
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	content := string(b)

	fmt.Println("\n#### Without decode AS IS ####\n" + content)

	// Decode CP1252 to unicode https://en.wikipedia.org/wiki/Windows-1252
	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(strings.NewReader(content))
	b, err = ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n######################################################################\n")
	fmt.Println("\n######## And now Ladies & Gentleman the Windows-1252 decoded in utf-8 : ########\n" + string(b))

	decodedContent := goutils.GetFileTextContent(filePath, "Windows1252")

	fmt.Println("\n######################################################################\n")
	fmt.Println("\n######## And now using getFileTextContent(filePath, 'Windows1252') : ########\n" + decodedContent)

}
