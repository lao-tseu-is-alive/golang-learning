package main

import (
	"bytes"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"log"
	"os/exec"
	"runtime"
)

func main() {

	cmd := "ls"
	arg := "-la"
	if runtime.GOOS == "windows" {
		cmd = "dir"
		arg = " /AH /AS /AA /C /Q"
	}

	process := exec.Command(cmd, arg)
	var out bytes.Buffer
	process.Stdout = &out
	err := process.Start()
	if err != nil {
		log.Fatalf("ERROR trying to start command : %s , error is : %v\n", cmd, err)
	}
	golog.Info("Process is running...")

	// Wait function will wait until the process end
	err = process.Wait()
	if err != nil {
		log.Fatalf("ERROR waiting for command : %s , error is : %v\n", cmd, err)
	}

	if process.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String())
	}
}
