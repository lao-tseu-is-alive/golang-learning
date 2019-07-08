package main

import (
	"github.com/lao-tseu-is-alive/golog"
	"regexp"
)

const baseString = "Mary was in love with Jack"

func main() {
	golog.Info("With this string :\n%s", baseString)
	const myRegex = "(Mary) (.+) (Jack)"
	golog.Warn("and this regex :\n%s", myRegex)
	regex := regexp.MustCompile(myRegex) // capture the 2 names and invert them
	out := regex.ReplaceAllString(baseString, "${3} ${2} ${1} ")
	golog.Info("after a : regex.ReplaceAllString(baseString, '${3} ${2} ${1} ')")
	golog.Info("You get has a result :\n%s", out)
}
