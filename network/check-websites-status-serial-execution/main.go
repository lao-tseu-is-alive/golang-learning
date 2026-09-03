package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var defaultUrls = []string{
	"https://www.google.com",
	"https://google.ch",
	"https://facebook.com",
	"https://stackoverflow.com",
	"https://golang.org",
	"https://amazon.com",
	"https://www.epfl.ch",
	"https://www.lausanne.ch",
	"https://www.cnn.com",
	"https://www.thissitewillprobablynotexist.com",
}

func timeTrack(start time.Time, name string, totalDataBytes *int) {
	elapsed := time.Since(start)
	if float64(elapsed/time.Millisecond) < float64(*totalDataBytes) {
		bytesPerMillisecond := float64(*totalDataBytes) / float64(elapsed/time.Millisecond)
		log.Printf("## %-50s took %s\t[%.2f bytes/ms]\t%9d bytes received\n", name, elapsed, bytesPerMillisecond, *totalDataBytes)

	} else {
		timePerByte := float64(elapsed/time.Millisecond) / float64(*totalDataBytes)
		log.Printf("## %-50s took %s\t[%.2f ms/bytes]\t%9d bytes received\n", name, elapsed, timePerByte, *totalDataBytes)
	}
}

func checkUrl(url string) (success bool, totalBytesReceived int) {
	numBytes := 0
	// here we defer the call and pass the reference to numBytes to get the final value of bytes received
	defer timeTrack(time.Now(), fmt.Sprintf("checkUrl('%s') ", url), &numBytes)
	res, err := http.Get(url)
	if err != nil {
		log.Printf("## 💥💥⚡ ERROR in checkUrl(url: %s ) doing http.Get(url) . Error was : %v\n", url, err)
		return false, numBytes
	}
	body, err := io.ReadAll(res.Body)
	numBytes = len(body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Printf("## 💥💥⚡ Response failed in checkUrl(url: %s ) with status code: %d, received %d bytes\n", url, res.StatusCode, numBytes)
	}
	if err != nil {
		log.Printf("## 💥💥⚡ ERROR in checkUrl(url: %s ) doing io.ReadAll(res.Body) . Error was : %v\n", url, err)
	}
	log.Printf("## Success in checkUrl(url: %s ) with status code: %d, received %d bytes\n", url, res.StatusCode, numBytes)
	return true, numBytes

}

func main() {
	totalBytesReceived := 0
	totalSuccess := 0
	totalFailures := 0
	defer timeTrack(time.Now(), "checkWebsitesStatusSerialExecution", &totalBytesReceived)
	for _, url := range defaultUrls {
		success, bytesReceived := checkUrl(url)
		if success {
			fmt.Printf("###### url : %s is OK ! ######\n", url)
			totalSuccess += 1
		} else {
			fmt.Printf("###### 💥💥 url : %s is NOT WORKING ! 💥💥 ######\n", url)
			totalFailures += 1
		}
		totalBytesReceived += bytesReceived
	}
	fmt.Printf("\nTotal bytes received : %d (%d success, %d failures)\n", totalBytesReceived, totalSuccess, totalFailures)

}
