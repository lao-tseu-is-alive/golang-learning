package main

import (
	"github.com/lao-tseu-is-alive/golog"
	"io/ioutil"
	"os"
)
import "bufio"

import "bytes"
import "fmt"

func readFileProgressive(filename string) {
	golog.Un(golog.Trace("readFileProgressive %s using reader", filename))
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the
	// file with reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

}

func readFileOneShot(filename string) {
	golog.Un(golog.Trace("readFileOneShot %s using ReadFile", filename))
	fContent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))
}

func main() {

	sampleFile := "data/JVerne_Aventures_du_capitaine_Hatteras.txt"
	golog.Info("### Read as reader ###")
	readFileProgressive(sampleFile)

	golog.Info("### ReadFile for smaller files only !###")
	readFileOneShot(sampleFile)

}
