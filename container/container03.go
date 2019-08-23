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
			run03()
		case "child":
			child()
		default:
			usage()
			panic(fmt.Sprintf("the command %q is not supported yet", os.Args[1]))
		}
	} else {
		usage()
	}
}

/*
 in this version, we still use namespace  by setting  CLONE_NEWUTS,
 but we also duplicate the process so that we can set some context on the child
*/
func run03() {
	fmt.Printf("PID[%d] About to run : %v \n ", os.Getpid(), os.Args[2:])
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// http://man7.org/linux/man-pages/man2/clone.2.html
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR doing cmd.Run()", err)
		os.Exit(1)
	}
}

/*
 now inside the child process pid=1 and hostname is already set to gocontainer
 but if you run ps you still see all the process
 so what if you really want to see only your container world
*/
func child() {
	fmt.Printf("PID[%d] About to run child : %v\n", os.Getpid(), os.Args[2:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// setting the hostname visible inside your go container
	if err := syscall.Sethostname([]byte("gocontainer03")); err != nil {
		fmt.Println("ERROR doing syscall.Sethostname()", err)
		os.Exit(1)
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR doing cmd.Run()", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("Usage : %s run bash ) \n", os.Args[0])
}
