package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Test runs a test suite against f.
func Test(f func(string) map[string]int) {
	ok := true
	for _, c := range testCases {
		got := f(c.in)
		if len(c.want) != len(got) {
			ok = false
		} else {
			for k := range c.want {
				if c.want[k] != got[k] {
					ok = false
				}
			}
		}
		if !ok {
			fmt.Printf("FAIL\n f(%q) =\n  %#v\n want:\n  %#v",
				c.in, got, c.want)
			break
		}
		fmt.Printf("PASS\n f(%q) = \n  %#v\n", c.in, got)
	}
}

var testCases = []struct {
	in   string
	want map[string]int
}{
	{"I am learning Go!", map[string]int{
		"I": 1, "am": 1, "learning": 1, "Go": 1,
	}},
	{"The quick brown fox jumped over the lazy dog.", map[string]int{
		"The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1,
		"over": 1, "the": 1, "lazy": 1, "dog": 1,
	}},
	{"I ate a donut. Then I ate another donut.", map[string]int{
		"I": 2, "ate": 2, "a": 1, "donut": 2, "Then": 1, "another": 1,
	}},
	{"A man a plan a canal panama.", map[string]int{
		"A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama": 1,
	}},
	{"Il l'a écoulé à l'avance, dans le château du plan d'eau de Viron-de-la-grande-motte.Malgré tous les risques!", map[string]int{
		"Il": 1, "Malgré": 1, "Viron-de-la-grande-motte": 1, "a": 1, "avance": 1, "château": 1, "d": 1, "dans": 1, "de": 1, "du": 1, "eau": 1, "l": 2, "le": 1, "les": 1, "plan": 1, "risques": 1, "tous": 1, "à": 1, "écoulé": 1,
	}},
}

func WordCount(s string) map[string]int {
	arrWord := strings.Fields(s)
	if len(strings.TrimSpace(s)) > 0 {
		// should match any utf8 letter including any diacritics
		// this regex will always match a rune like : à, regardless of how it is encoded !
		// https://www.regular-expressions.info/unicode.html
		reWords := regexp.MustCompile("(\\p{L}(\\p{M}|-)*)+")
		arrWord = reWords.FindAllString(s, -1)
	} else {
		panic("String should not be spaces only !")
	}
	res := make(map[string]int)
	for _, v := range arrWord {
		//fmt.Printf("k[%v] = %v\n",k,v)
		val, ok := res[v]
		if ok {
			res[v] = val + 1
		} else {
			res[v] = 1
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	Test(WordCount)
}
