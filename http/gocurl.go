package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	//"time"
)

func main() {
	const defaultUrl = "https://postman-echo.com/time/now"
	var url = defaultUrl
	flag.StringVar(&url, "url", defaultUrl, "url to query")
	flag.Parse()

	/*tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	*/
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}
