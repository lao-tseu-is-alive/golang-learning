package main

import (
	"fmt"
	"os"
	"os/exec"
)

/***
	Inspired by the great presentation from Liz Rice "Containers from Scratch"
	https://www.youtube.com/watch?v=8fi7uSYlOdc
	the idea is to create our own go program that will allow doing the same as :
	docker run -it --rm ubuntu bash
	docker run --interactive --tty ubuntu bash
	and then apply constrained resources :
	https://docs.docker.com/config/containers/resource_constraints/
	docker run -it -m64m --rm ubuntu bash
 ***/
func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "run":
			run01()
		default:
			usage()
			panic(fmt.Sprintf("the command %q is not supported yet", os.Args[1]))
		}
	} else {
		usage()
	}
}

/*
 this one is as simple as it can be, no container isolation at all
 if you change hostname it will survive after exit : this is BAD !
 we are simply forking a new bash that's all
*/
func run01() {
	fmt.Printf("About to run : %v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR doing cmd.Run()", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("Usage : %s run bash --rcfile <(cat ~/.bashrc; echo 'PS1=\"gocontainer01>\"')  ... \n", os.Args[0])
}
