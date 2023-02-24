package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

/*
	allow you to pipe a text into this unix filter program and get back the list of utf8 words one by line

cat your_utf8_text.txt | getwords --min-length=2 --show-line-number=0
by combining with classical unix filters you can get a ordered list of top ten most used words in the text !
cat your_utf8_text.txt | getwords --min-length=3 --show-line-number=0 -convert-to-case=l |sort |uniq -c |sort -nr | head
*/
func main() {
	args := os.Args
	app := args[0]
	sc := bufio.NewScanner(os.Stdin)
	minWordLength := flag.Int("min-length", 1, "minimum length of words to keep (default is 1)")
	showLineNumber := flag.Bool("show-line-number", false, "allow display of prefix with line number")
	convertToCase := flag.String("convert-to-case", "", "convert case of all words. Can be one of (U,u,upper or L,l,lower)")
	flag.Parse()
	count := 1

	for sc.Scan() {
		if len(strings.TrimSpace(sc.Text())) > 0 {
			// should match any utf8 letter including any diacritics
			// this regex will always match a rune like : à, regardless of how it is encoded !
			// https://www.regular-expressions.info/unicode.html
			reWords := regexp.MustCompile("(\\p{L}\\p{M}*)+")
			words := reWords.FindAllString(sc.Text(), -1)
			onlyBiggerWords := []string{}
			for _, word := range words {
				if utf8.RuneCountInString(strings.TrimSpace(word)) >= *minWordLength {
					switch *convertToCase {
					case "upper", "u", "U":
						onlyBiggerWords = append(onlyBiggerWords, strings.ToUpper(word))
					case "lower", "l", "L":
						onlyBiggerWords = append(onlyBiggerWords, strings.ToLower(word))
					default:
						onlyBiggerWords = append(onlyBiggerWords, word)
					}
				}
			}
			if len(onlyBiggerWords) > 0 {
				if *showLineNumber {
					wList := strings.Join(onlyBiggerWords, " ")
					fmt.Printf("%d : %s \n", count, wList)
				} else {
					wList := strings.Join(onlyBiggerWords, "\n")
					fmt.Printf("%s\n", wList)
				}
			}
			count += 1
			//fmt.Println(wList)
		} else {
			if count < 2 {
				//TODO find a solution to display help when no data is present without polutin output if first line is empty
				golog.Warn("# %s should be used in a pipe like this : cat your_utf8_text.txt | getwords --min-length=2 --show-line-number=0", app)
			}
		}
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
