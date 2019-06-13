package main

import (
	"github.com/lao-tseu-is-alive/golog"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func main() {
	exeFullPath, err := os.Executable()
	if err != nil {
		log.Fatalln("Error getting Executable Path ! %v", err)
	}
	golog.Info("Full path of executable : %s", exeFullPath)
	exeDirectory := filepath.Dir(exeFullPath)
	golog.Info("Directory of executable : %s", exeDirectory)
	// Use EvalSymlinks to get the real path.
	realPath, err := filepath.EvalSymlinks(exeDirectory)
	if err != nil {
		log.Fatalln("Error getting real Path ! %v", err)
	}
	golog.Info("Symlink evaluated  dir  : %s", realPath)

	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error getting working directory ! %v", err)
	}
	golog.Info("Current working directory: %s", workingDirectory)
	// detecting OS at runtime next line gives you linux or windows etc..
	golog.Warn("Current system is %s", runtime.GOOS)
	// this is how you can get the PATH separator
	golog.Info("Directory separator is :%q ", os.PathSeparator)
	configDir := path.Join(exeDirectory, "config")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		golog.Err("the config subdirectory does not exist here : %s", configDir)
	} else {
		golog.Info("Found config directory %s", configDir)
	}
	parentDataDir := path.Join(exeDirectory, "../data")
	if _, err := os.Stat(parentDataDir); os.IsNotExist(err) {
		golog.Err("the data directory does not exist here : %s", parentDataDir)
	} else {
		golog.Info("Found config directory %s", parentDataDir)
	}
}
