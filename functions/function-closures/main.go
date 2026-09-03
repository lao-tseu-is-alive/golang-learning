package main

import "fmt"

func adder(name string) func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Printf("IN %s(x:%v)\tsum :%v\n", name, x, sum)
		sum += x
		return sum
	}
}

/*
Go functions may be closures. A closure is a function value that references
variables from outside its body. The function may access and assign to the referenced variables;
in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.
*/

func main() {
	pos, neg := adder("adder"), adder("neg")
	for i := 0; i < 10; i++ {
		fmt.Printf("############ i=%d ############\n", i)
		fmt.Printf("pos(i) = %v\t neg(-2*i) = %v\n",
			pos(i),
			neg(-2*i),
		)
	}
}
