package main

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	app := args[0]
	const verboseHelp = "allow to get more verbose feedback."
	var isVerbose = false
	flag.BoolVar(&isVerbose, "verbose", false, verboseHelp)
	flag.BoolVar(&isVerbose, "v", false, verboseHelp+"(shorthand)")
	flag.Parse()
	numbers := args[1:]
	if len(args) > 1 {
		if isVerbose {
			golog.Info(fmt.Sprintf("%s was called whith those arguments :", app))
			// The rest of the arguments could be obtained by omitting the first argument.

			for idx, arg := range numbers {
				golog.Info(fmt.Sprintf("Arg %d = %s \n", idx, arg))
			}
		}

		if len(numbers) < 1 {
			golog.Warn(fmt.Sprintf("I did not receive any numbers, so result is obviously : %s", "zero"))
		} else {
			var sum int64 = 0
			for idx, arg := range numbers {
				if isVerbose {
					golog.Warn(fmt.Sprintf("(sum = %d) will try to add arg[%d]=%s ", sum, idx, arg))
				}
				num, err := strconv.ParseInt(arg, 10, 64)
				if err != nil {
					golog.Err(fmt.Sprintf("Array[%d]=%s IS NOT A VALID INTEGER bypassing this one !", idx, arg))
				} else {
					sum += num
				}
			}
			if isVerbose {
				golog.Info(fmt.Sprintf("Array sum total is = %d", sum))
			}
			fmt.Println(sum)
		}
	} else {
		flag.PrintDefaults()
		fmt.Printf("Usage : %s 1 2 4 5\nwill return the sum of those numbers : 12 (=1+2+4+5)\n", app)
	}

}
