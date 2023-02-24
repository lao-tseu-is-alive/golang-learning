package main

import (
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/unicode/runenames"
	"strings"
	"unicode"
	"unicode/utf8"
)

func getUnicodeInfo(c rune) string {
	res := "is "
	if unicode.IsControl(c) {
		res = res + "control,"
	}
	if unicode.IsDigit(c) {
		res = res + "digit,"
	}
	if unicode.IsGraphic(c) {
		res = res + "graphic,"
	}

	if unicode.IsMark(c) {
		res = res + "mark,"
	}
	if unicode.IsPrint(c) {
		res = res + "printable,"
	}
	if !unicode.IsPrint(c) {
		res = res + "not printable,"
	}
	if unicode.IsPunct(c) {
		res = res + "punct,"
	}
	if unicode.IsSpace(c) {
		res = res + "space,"
	}
	if unicode.IsSymbol(c) {
		res = res + "symbol,"
	}
	if unicode.IsTitle(c) {
		res = res + "title case,"
	}
	if unicode.IsNumber(c) {
		res = res + "number,"
	}
	if unicode.IsLetter(c) {
		res = res + "letter,"
	}
	if unicode.IsUpper(c) {
		res = res + "upper case,"
	}
	if unicode.IsLower(c) {
		res = res + "lower case,"
	}

	return res
}

// normalize will give the normal form version of a given unicode string
// https://go.dev/blog/normalization
func normalize(s string, toLower bool) string {
	t := transform.Chain(norm.NFKD, runes.Remove(runes.In(unicode.Mn)), norm.NFKD)
	res, _, _ := transform.String(t, s)
	if toLower {
		return strings.ToLower(res)
	}
	return res
}

/*
	runesUnicode will allow you to see every rune in the string and the underlying bytes

example output of single run :
cgil@pulsar2021:~/cgdev/golang/golang-learning/strings$ go run runesUnicode.go
go run runesUnicode.go
character '🔥' [U+1F525] FIRE, 	normalized:'🔥'  starts at byte 0 and is graphic,printable,symbol,
0: 	 240 [F0]
1: 	 159 [9F]
2: 	 148 [94]
3: 	 165 [A5]
character 'ℂ' [U+2102] DOUBLE-STRUCK CAPITAL C, 	normalized:'c'  starts at byte 4 and is graphic,printable,letter,upper case,
0: 	 226 [E2]
1: 	 132 [84]
2: 	 130 [82]
character '𝕒' [U+1D552] MATHEMATICAL DOUBLE-STRUCK SMALL A, 	normalized:'a'  starts at byte 7 and is graphic,printable,letter,lower case,
0: 	 240 [F0]
1: 	 157 [9D]
2: 	 149 [95]
3: 	 146 [92]
the string '🔥ℂ𝕒' contains 11 bytes, and 3 runes
*/
func main() {
	sampleText := "🔥ℂ𝕒𝕣𝕝𝕠𝕤❤️☯１２３C͎A͓̽RL҉O҉S҉ ҉҉̾ \v\a💌 💕💞💓💗💖💘💝💟💜💛🧡️💚💙💔🌲️🌳️🌴️🌍️🏔️☀️🌞️⭐️💥"
	//sampleText := "🔥ℂ𝕒"
	indexByte := 0
	for pos, rune := range sampleText {
		normalized := normalize(string(rune), true)
		fmt.Printf("character '%q'\t[%U] %-30q, \tnormalized:'%s'  starts at byte %d and %s\n",
			rune, rune, runenames.Name(rune), normalized, pos, getUnicodeInfo(rune))
		for i, byteVal := range []byte(string(rune)) {
			fmt.Printf("%5d%3d: \t %d [%X]\n", indexByte, i, byteVal, byteVal)
			indexByte += 1
		}
	}

	fmt.Printf("the string '%s' contains %d bytes, and %d runes\n",
		sampleText, len(sampleText), utf8.RuneCountInString(sampleText))

}
