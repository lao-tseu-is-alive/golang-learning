package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%s\n", &name)

	var age int
	fmt.Println("What is your age?")
	fmt.Scanf("%d\n", &age)

	fmt.Printf("Hello %s, nice to meet you ! \n", name)

	if age < 10 {
		fmt.Printf("%d years only, and already playing with golang Waouuh ! \n", age)
	}

}
