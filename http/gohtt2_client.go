package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://localhost:8000"

var httpVersion = flag.Int("version", 2, "HTTP version")

/*
code from article : https://developpaper.com/the-http-2-server-function-of-go-language-and-the-use-of-client/
*/
func main() {
	flag.Parse()
	client := &http.Client{}
	// Create a pool with the server certificate since it is not signed
	// by a known CA
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	switch *httpVersion {
	case 1:
		client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	case 2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	}
	// Perform the request
	resp, err := client.Get(url)

	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	fmt.Printf("Got response %d: %s %s\n", resp.StatusCode, resp.Proto, string(body))
}
