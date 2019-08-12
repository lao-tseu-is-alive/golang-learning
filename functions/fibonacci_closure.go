package main

import "fmt"

/* fibonacci is a function that returns
 a function that returns an int.
 0, 1, 1, 2, 3, 5, 8, 13 and 21
https://en.wikipedia.org/wiki/Fibonacci_number
*/
func fibonacci() func() int {
	x, y, sum := 0, 1, 0
	return func() int {
		// fmt.Printf("f(x:%v, y:%v)\tsum :%v\n", x, y, sum)
		sum, x, y = x, y, x+y
		// fmt.Printf("f(x:%v, y:%v)\tsum :%v\n", x, y, sum)
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 11; i++ {
		//fmt.Printf("### %v : ###\n", i)
		fmt.Printf("%v, ", f())
	}
}
