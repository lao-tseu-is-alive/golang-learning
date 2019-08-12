package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// 1, 1, 2, 3, 5, 8, 13 and 21
func fibonacci() func() int {
	x := 0
	y := 1
	sum := 0
	return func() int {
		fmt.Printf("f(x:%v, y:%v)\tsum :%v\n", x, y, sum)
		sum, x, y = x, y, x+y
		fmt.Printf("f(x:%v, y:%v)\tsum :%v\n", x, y, sum)
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("### %v : ###\n", i)
		fmt.Printf("%v\n", f())
	}
}
