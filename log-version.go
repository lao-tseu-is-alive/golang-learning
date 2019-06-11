package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

const app = "log-version"

// because Info and Trace are defined as public func in log.go in the same directory
// to run/build  you need to run in the shell :
// go run log-version.go log.go
func main() {
	Trace("BEGIN in main()")
	defer Trace("END of main()")
	const info = `The binary was build by Go version : %s`

	log.Printf("Application '%s' is starting\n", app)
	log.Printf(info, runtime.Version())
	t := time.Now()
	timeString := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d.%03d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/100000)
	Err(fmt.Sprintf("%s something went wrong at ", timeString))
	Info("just a simple information message to send to log")
}
