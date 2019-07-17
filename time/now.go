package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	fmt.Println(today)                           // 2019-07-16 14:42:44.172473354 +0200 CEST m=+0.000048118
	fmt.Println(today.Format("02-Jan-2006"))     // 16-Jul-2019
	fmt.Println(today.Format("20060102_150405")) // 20190716_144244
}
