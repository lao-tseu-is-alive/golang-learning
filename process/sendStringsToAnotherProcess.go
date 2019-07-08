package main

import (
	"bufio"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"io"
	"os/exec"
	"time"
)

func main() {
	cmd := []string{"go", "run", "echoStdin.go"}

	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	// For debugging purposes we watch the
	// output of the executed process
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Program says:" + s.Text())
		}
	}()

	// Start the process
	golog.DoItOrDie(proc.Start(), "starting process")

	// Now the following lines
	// are written to child
	// process standard input
	fmt.Println("Writing input")
	io.WriteString(stdin, "Hello\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "is awesome\n")

	time.Sleep(time.Second * 2)

	proc.Process.Kill()

}
