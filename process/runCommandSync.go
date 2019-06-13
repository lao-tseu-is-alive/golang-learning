package main

import (
	"bytes"
	"fmt"
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
	err := process.Run()
	if err != nil {
		log.Fatalf("ERROR trying to run command : %s , error is : %v\n", cmd, err)
	}
	fmt.Println("Process is now finished, here are the results :")
	fmt.Println(out.String())
}
