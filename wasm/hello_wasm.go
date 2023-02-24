package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello WebAssembly ! From go version : %v\n", runtime.Version())
	fmt.Println("More info : https://github.com/golang/go/wiki/WebAssembly")
	fmt.Println("COMPILED FOR WASM WITH : \nGOOS=js GOARCH=wasm go build -o hello.wasm hello_wasm.go")

}
