package main

import (
	"context"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"os/exec"
	"runtime"
	"time"
)

func trace() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("%s,:%d %s\n", frame.File, frame.Line, frame.Function)
}

func innerFunction() {
	defer golog.Un(golog.Trace("innerFunction"))
	golog.PrintCallStack()
}
func outerFunction() {
	defer golog.Un(golog.Trace("outerFunction"))
	innerFunction()
}

func main() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	// The provided context is used to kill the myCmd (by calling os.Process.Kill)
	// if the context becomes done before the command completes on its own.
	// https://golang.org/pkg/os/exec/#CommandContext
	myCmd := exec.CommandContext(ctx, cmd, "1") // chilD process should wait 1 sec
	golog.DoItOrDie(myCmd.Start(), "Doing myCmd.Start(), cmd: %s", cmd)

	golog.Info("PID (after Start, before Wait): %d\n", myCmd.Process.Pid)
	// myCmd.ProcessState is nil until the myCmd process as finish.
	golog.Info("Process state for running child myCmd: %v\n", myCmd.ProcessState)
	// This will fail after 500 milliseconds. The child myCmd with 1 second sleep will be interrupted.
	// Wait function will wait until the myCmd ends (with success or because an error occurs)
	golog.DoItOrDie(myCmd.Wait(), "Doing myCmd.Wait(), Child PID: %d", myCmd.Process.Pid)

	// After the myCmd terminates the *os.ProcessState contains simple information about the child myCmd run
	golog.Info("PID (after Wait): %d\n", myCmd.ProcessState.Pid())
	golog.Info("Process took: %dms\n", myCmd.ProcessState.SystemTime()/time.Microsecond)
	golog.Info("Exited sucessfuly ? : %t\n", myCmd.ProcessState.Success())

	// The PID could be obtained for the running myCmd
	fmt.Printf("PID of running myCmd: %d\n\n", myCmd.Process.Pid)
	// test callStack
	outerFunction()
}
