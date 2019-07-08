package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

func main() {
	/*
		Replacer structure also has the WriteString method.
		This method will write to the given writer with all replacements defined in Replacer.
		The main purpose of this type is its reusability.
		It can replace multiple strings at once and it is safe for concurrent use !
	*/
	replacer := strings.NewReplacer(
		"lamb", "wolf",
		"Mary", "Jack")
	out := replacer.Replace(refString)
	fmt.Println(out)
}
