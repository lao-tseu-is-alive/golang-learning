package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}    // has type Vertex
	v  = Vertex{91, 92}  // has type
	v2 = Vertex{X: 1}    // Y:0 is implicit
	v3 = Vertex{}        // X:0 and Y:0
	p  = &Vertex{11, 21} // has type *Vertex
	p1 = &v
)

func main() {
	p1.X = 5
	fmt.Println(v1, p, p1, v2, v3)
	fmt.Printf("v1 : %T\t%v\n", v1, v1)
	fmt.Printf("v  : %T\t%v\n", v, v)
	fmt.Printf("p1 : %T\t%v\n", p1, p1)
}
