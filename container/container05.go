package main

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

const ROOTFS = "/home/cgil/go/src/github.com/lao-tseu-is-alive/golang-learning/container/alpinelinux_2.10.4-r2/ROOTFS"
const CGName = "golang"

/*
**

		Inspired by the great presentation from Liz Rice "Containers from Scratch"
		https://www.youtube.com/watch?v=8fi7uSYlOdc
		the idea is to create our own go program that will allow doing the same as :
		docker run -it --rm ubuntu bash
		docker run --interactive --tty ubuntu bash
		and then apply constrained resources :
		https://docs.docker.com/config/containers/resource_constraints/
		docker run -it -m64m --rm ubuntu bash
		Sécurité des infrastructures de virtualisation : https://tim.siosm.fr/cours/static/2018-2019/infra-virt-sec-print.pdf
	 **
*/
func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "run":
			run05()
		case "child":
			child05()
		default:
			usage()
			panic(fmt.Sprintf("the command %q is not supported yet", os.Args[1]))
		}
	} else {
		usage()
	}
}

/*
	 in this version, we use CLONE_NEWNS to limit visibility of mount outside of the container
	 and we use the CGROUPS to limit the number of the process
	 but we also duplicate the process so that we can set some context on the child
	 how to protect our self against a simple shell bomb  :() { : | : &};:
	 in docker you can limit the number of process by container with --pids-limit
		docker run -it --rm --pids-limit=3 --memory=25m ubuntu bash
		https://docs.docker.com/engine/reference/commandline/run/#options
		https://docs.docker.com/config/containers/resource_constraints/
		https://www.serverlab.ca/tutorials/containers/docker/how-to-limit-memory-and-cpu-for-docker-containers/
		https://hostadvice.com/how-to/how-to-limit-a-docker-containers-resources-on-ubuntu-18-04/
*/
func run05() {
	fmt.Printf("PID[%d] About to run : %v \n ", os.Getpid(), os.Args[2:])
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// http://man7.org/linux/man-pages/man2/clone.2.html
	// https://golang.org/pkg/syscall/#SysProcAttr
	cmd.SysProcAttr = &syscall.SysProcAttr{

		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS, // to avoiding share mount with the host
	}

	if err := cmd.Run(); err != nil {
		golog.Err("ERROR doing cmd.Run()", err)
		os.Exit(1)
	}
}

/*
now inside the child process after pid=1 and specific hostname for the gocontainer
we mount a specific filesystem and use chroot to isolate this  process
so what if you really want to see only your container world
*/
func child05() {
	fmt.Printf("PID[%d] About to run child : %v\n", os.Getpid(), os.Args[2:])

	cgMaxProcess(20)

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// setting the hostname visible inside your go container
	if err := syscall.Sethostname([]byte("gocontainer05")); err != nil {
		golog.Err("ERROR doing syscall.Sethostname()", err)
		os.Exit(1)
	}
	if err := syscall.Chroot(ROOTFS); err != nil {
		golog.Err("ERROR doing syscall.Chroot(ROOTFS)", err)
		os.Exit(1)
	}
	if err := syscall.Chdir("/"); err != nil {
		golog.Err("ERROR doing syscall.Chdir(/)", err)
		os.Exit(1)
	}
	if err := syscall.Mount("proc", "proc", "proc", 0, ""); err != nil {
		golog.Err("ERROR doing syscall.Mount(proc ...)", err)
		os.Exit(1)
	}

	if err := cmd.Run(); err != nil {
		golog.Err("ERROR doing cmd.Run()", err)
		os.Exit(1)
	}

	if err := syscall.Unmount("/proc", 0); err != nil {
		golog.Err("ERROR doing syscall.Unmount(proc)", err)
		os.Exit(1)
	}
}

/*
we use CGROUPS limit the number of process like the --pids-limit=3 in docker
*/
func cgMaxProcess(maxProcess int) {
	cgroups := "/sys/fs/cgroup/"
	fullPidsPath := filepath.Join(cgroups, "pids", CGName)
	if _, err := os.Stat(fullPidsPath); os.IsNotExist(err) {
		err := os.Mkdir(fullPidsPath, 0755)
		if err != nil {
			golog.Err("ERROR in os.Mkdir [%s] error is : %v\n", fullPidsPath, err)
			os.Exit(1)
		}
	}

	err := ioutil.WriteFile(filepath.Join(fullPidsPath, "pids.max"), []byte(strconv.Itoa(maxProcess)), 0700)
	if err != nil {
		golog.Err("ERROR in cgMaxProcess writing pids.max", err)
		os.Exit(1)
	}
	// Removes the new cgroup in place after the container exits
	err = ioutil.WriteFile(filepath.Join(fullPidsPath, "notify_on_release"), []byte("1"), 0700)
	if err != nil {
		golog.Err("ERROR in cgMaxProcess writing notify_on_release", err)
		os.Exit(1)
	}
	err = ioutil.WriteFile(filepath.Join(fullPidsPath, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700)
	if err != nil {
		golog.Err("ERROR in cgMaxProcess writing cgroup.procs", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Printf("Usage : %s run bash -l) \n", os.Args[0])
}
