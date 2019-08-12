package main

import "fmt"

type VertexInt struct {
	X int
	Y int
}

var (
	v1 = VertexInt{1, 2}    // has type VertexInt
	v  = VertexInt{91, 92}  // has type
	v2 = VertexInt{X: 1}    // Y:0 is implicit
	v3 = VertexInt{}        // X:0 and Y:0
	p  = &VertexInt{11, 21} // has type *VertexInt
	p1 = &v
)

func main() {
	p1.X = 5
	fmt.Println(v1, p, p1, v2, v3)
	fmt.Printf("v1 : %T\t%v\n", v1, v1)
	fmt.Printf("v  : %T\t%v\n", v, v)
	fmt.Printf("p1 : %T\t%v\n", p1, p1)
}
