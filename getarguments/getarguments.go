package main

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"os"
	"strconv"
	"strings"
)

// Custom type need to implement
// flag.Value interface to be able to
// use it in flag.Var function.
type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

func main() {
	args := os.Args
	app := args[0]

	// Extracting flag values with methods returning pointers
	retry := flag.Int("retry", -1, "Defines max retry count")

	// Read the flag using the XXXVar function.
	// In this case the variable must be defined prior to the flag.
	var infoPrefix string
	flag.StringVar(&infoPrefix, "prefix", "", "prefix string")

	var arr ArrayValue
	flag.Var(&arr, "array", "Input array of integers separated by , to iterate through.")

	// Execute the flag.Parse function, to read the flags to defined variables.
	// Without this call the flag variables remain empty.
	flag.Parse()

	golog.Info(fmt.Sprintf("%s was called whith those arguments :", app))
	// The rest of the arguments could be obtained
	// by omitting the first argument.
	otherArgs := args[1:]

	fmt.Println(otherArgs)

	for idx, arg := range otherArgs {
		golog.Info(fmt.Sprintf("Arg %d = %s \n", idx, arg))
	}

	if *retry == -1 {
		golog.Err("you should provide a valid parameter -retry as number of retries")
	} else {
		golog.Warn(fmt.Sprintf("retry value is : %d", *retry))
	}
	if len(infoPrefix) > 0 {
		golog.Warn(fmt.Sprintf("prefix value is : %s", infoPrefix))
	}

	if len(arr) > 0 {
		var sum int64 = 0
		for idx, arg := range arr {
			golog.Warn(fmt.Sprintf("Array[%d]=%s", idx, arg))
			num, err := strconv.ParseInt(arg, 10, 64)
			if err != nil {
				golog.Err(fmt.Sprintf("Array[%d]=%s IS NOT A VALID INTEGER bypassing this one !", idx, arg))
			} else {
				sum += num
			}
		}
		golog.Info(fmt.Sprintf("Array sum total is = %d", sum))
	}

}
