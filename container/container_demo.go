package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*
	inspired by the great presentation from Liz Rice "Containers from Scratch"
	https://www.youtube.com/watch?v=8fi7uSYlOdc

https://www.technoarete.org/common_abstract/pdf/IJERCSE/v5/i3/Ext_97135.pdf
https://www.infoq.com/articles/build-a-container-golang/
*/
func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "run":
			parent()
		case "child":
			child()
		default:
			panic("wat should I do")
		}
	} else {
		fmt.Printf("Usage : %s run yourCommand param1 ... ", os.Args[0])
	}
}

func parent() {
	fmt.Printf("About to run : %v\n", os.Args[2:])
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR doing cmd.Run()", err)
		os.Exit(1)
	}
}

func child() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
