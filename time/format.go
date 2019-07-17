package main

import (
	"fmt"
	"time"
)

const timeFormatCheatsheet = `
#Date and Time Options Cheatsheet :
You can use this table below as a quick cheatsheet for the different date and time options and the layout params you need to use.

Type	Options
Year	06, 2006
Month	01, 1, Jan, January
Day		02, 2, _2, (width two, right justified)
Weekday	Mon, Monday
Hours	03, 3, 15
Minutes	04, 4
Seconds	05, 5
ms μs ns	.000, .000000, .000000000
ms μs ns	.999, .999999, .999999999 (trailing zeros removed)
am/pm	PM, pm
Timezone	MST
Offset	-0700, -07, -07:00, Z0700, Z07:00
`

func main() {
	tTime := time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)

	// The formatting is done
	// with use of reference value
	// Jan 2 15:04:05 2006 MST
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))

	fmt.Printf("The time is: %s\n", tTime.Format("15:04"))

	//The predefined formats could
	// be used
	fmt.Printf("The time is: %s\n", tTime.Format(time.RFC1123))

	// The formatting supports space padding
	//only for days in Go version 1.9.2
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/_2"))

	// The zero padding is done by adding 0
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/01/02"))

	//The fraction with leading zeros use 0s
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.00"))

	//The fraction without leading zeros use 9s
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.999"))

	// Append format appends the formatted time to given
	// buffer
	fmt.Println(string(tTime.AppendFormat([]byte("The time 	is up: "), "03:04PM")))
	fmt.Printf("%s", timeFormatCheatsheet)
}
