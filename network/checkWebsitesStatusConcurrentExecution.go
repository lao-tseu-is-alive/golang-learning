package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

var urls = []string{}

type urlCheckInfo struct {
	url           string
	success       bool
	status        int
	receivedBytes int
}

func timeMonitor(start time.Time, name string, totalDataBytes *int) {
	elapsed := time.Since(start)
	if float64(elapsed/time.Millisecond) < float64(*totalDataBytes) {
		bytesPerMillisecond := float64(*totalDataBytes) / float64(elapsed/time.Millisecond)
		log.Printf("## %-50s took %s\t[%.2f bytes/ms]\t%9d bytes received\n", name, elapsed, bytesPerMillisecond, *totalDataBytes)

	} else {
		timePerByte := float64(elapsed/time.Millisecond) / float64(*totalDataBytes)
		log.Printf("## %-50s took %s\t[%.2f ms/bytes]\t%9d bytes received\n", name, elapsed, timePerByte, *totalDataBytes)
	}
}

func checkUrlConcurrent(url string, totalBytesReceived chan *urlCheckInfo) {
	numBytes := 0
	urlInfo := urlCheckInfo{
		url:           url,
		success:       false,
		status:        0,
		receivedBytes: 0,
	}
	// here we defer the call and pass the reference to numBytes to get the final value of bytes received
	defer timeMonitor(time.Now(), fmt.Sprintf("checkUrl('%s') ", url), &numBytes)
	res, err := http.Get(url)
	if err != nil {
		log.Printf("## 💥💥⚡ ERROR in checkUrl(url: %s ) doing http.Get(url) . Error was : %v\n", url, err)
		totalBytesReceived <- &urlInfo
		return
	}
	body, err := io.ReadAll(res.Body)
	numBytes = len(body)
	urlInfo.receivedBytes = numBytes
	urlInfo.status = res.StatusCode
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Printf("## 💥💥⚡ Response failed in checkUrl(url: %s ) with status code: %d, received %d bytes\n", url, res.StatusCode, numBytes)
	}
	if err != nil {
		log.Printf("## 💥💥⚡ ERROR in checkUrl(url: %s ) doing io.ReadAll(res.Body) . Error was : %v\n", url, err)
	} else {
		log.Printf("## Success in checkUrl(url: %s ) with status code: %d, received %d bytes\n", url, res.StatusCode, numBytes)
	}
	urlInfo.success = true
	totalBytesReceived <- &urlInfo

}

func main() {
	fmt.Println("Go version", runtime.Version())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	args := os.Args
	app := args[0]
	// The rest of the arguments could be obtained
	// by omitting the first argument.
	otherArgs := args[1:]

	var urlFile string
	flag.StringVar(&urlFile, "urlFile", "urls.txt", "file containing the list of url to check")

	// Execute the flag.Parse function, to read the flags to defined variables.
	// Without this call the flag variables remain empty.
	flag.Parse()

	log.Printf("%s called whith arguments %d arguments", app, len(otherArgs))
	log.Printf("urlFile argument is  : %s", urlFile)

	content, err := os.ReadFile(urlFile)
	if err != nil {
		log.Fatalf("ERROR trying os.ReadFile(%s) error was: %v", urlFile, err)
	}
	urls := strings.Split(string(content), "\n")

	urlChannel := make(chan *urlCheckInfo)
	urlInfo := &urlCheckInfo{}
	totalBytesReceived := 0
	totalSuccess := 0
	totalFailures := 0
	defer timeMonitor(time.Now(), "checkWebsitesStatusConcurrentExecution", &totalBytesReceived)

	for _, url := range urls {
		go checkUrlConcurrent(url, urlChannel)
	}

	for _ = range urls {
		urlInfo = <-urlChannel
		log.Printf("**** urlInfo received : %v+ ****\n", urlInfo)
		if urlInfo.success {
			//fmt.Printf("###### url : %s is OK ! ######\n", urlInfo.url)
			totalSuccess += 1
		} else {
			fmt.Printf("###### 💥💥 url : %s is NOT WORKING ! 💥💥 ######\n", urlInfo.url)
			totalFailures += 1
		}
		totalBytesReceived += urlInfo.receivedBytes
	}

	fmt.Printf("\nTotal bytes received : %d (%d success, %d failures)\n", totalBytesReceived, totalSuccess, totalFailures)

}
