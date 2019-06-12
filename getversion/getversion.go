package main

import (
	"github.com/lao-tseu-is-alive/golog"
	"log"
	"runtime"
)

const app = "getversion"

// because Info and Trace are defined as public func in log.go in the same directory
// to run/build  you need to run in the shell :
// go run getversion.go log.go
func main() {
	golog.Trace("BEGIN in main()")
	defer golog.Trace("END of main()")
	const info = `The binary was build by Go version : %s`

	log.Printf("Application '%s' is starting\n", app)
	log.Printf(info, runtime.Version())
	golog.Err("something went terribly wrong here !")
	golog.Warn("this is a warning !")
	golog.Info("just a simple information message to send to log")
}
