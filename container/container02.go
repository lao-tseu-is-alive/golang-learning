package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
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
			run02()
		default:
			usage()
			panic(fmt.Sprintf("the command %q is not supported yet", os.Args[1]))
		}
	} else {
		usage()
	}
}

/*
 in this version, we begin using namespace  by setting  CLONE_NEWUTS
	If CLONE_NEWUTS is set, the process is created in a new UTS namespace,
 so now if you change hostname it will have no effect outside the process and after exit
 this is better... btu what about setting the hostname before we run the client command that's the next step
*/
func run02() {
	fmt.Printf("About to run : %v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// http://man7.org/linux/man-pages/man2/clone.2.html
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR doing cmd.Run()", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("Usage : %s run bash --rcfile <(cat ~/.bashrc; echo 'PS1=\"gocontainer02>\"') \n", os.Args[0])
}
