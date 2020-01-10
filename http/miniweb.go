package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var helpString = `miniweb : is a basic http server written in Golang
	you can configure via ENV variables or classic Parameter:
		-webPort or WEB_PORT : the Tcp port the server will listen to http request
		-webRootDir or WEB_ROOT_DIR : the path of the directory you want to serve via web
	to serve web  http request on port 8080 all the files in directory /var/www run one of this :
		go run miniweb.go -webPort 8888 -webRootDir /var/www
		WEB_ROOT_DIR=/var/www WEB_PORT=8888 go run miniweb.go`

func main() {
	var webPort int
	var webRootDir string
	// beginning of parameters handling
	val, exist := os.LookupEnv("WEB_PORT")
	if !exist {
		flag.IntVar(&webPort, "webPort", 8080, "Tcp port the server will listen to http request")
	} else {
		log.Printf("Using ENV variable WEB_PORT : [%v] to listen \n", val)
		webPort, _ = strconv.Atoi(val)
	}
	val, exist = os.LookupEnv("WEB_ROOT_DIR")
	if !exist {
		flag.StringVar(&webRootDir, "webRootDir", ".", "Path of the directory you want to serve via web")
	} else {
		log.Printf("Using ENV variable WEB_ROOT_DIR [%v] as root web path \n", val)
		webRootDir = val
	}
	helpPtr := flag.Bool("help", false, "Get help on options to use this program")
	flag.Parse()
	if *helpPtr == true {
		fmt.Println(helpString)
		os.Exit(0)
	}
	// end of parameters handling

	webRootDir, err := filepath.Abs(webRootDir)
	if err != nil {
		log.Fatalf("Problem getting absolute path of directory: %s\nError:\n%v\n", webRootDir, err)
	}
	if _, err := os.Stat(webRootDir); os.IsNotExist(err) {
		log.Fatalf("The webRootDir parameter is wrong, %s is not a valid directory\nError:\n%v\n", webRootDir, err)
	}

	log.Printf("MINIWEB STARTED : listening on port %d\nServing files from: %s\n", webPort, webRootDir)

	http.Handle("/", http.FileServer(http.Dir(webRootDir)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", webPort), nil))
}
