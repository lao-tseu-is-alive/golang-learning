package main

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"github.com/lao-tseu-is-alive/goutils"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type htmlData = struct {
	TitlePage string
	Headers   []string
}

const (
	defaultHost        = "localhost"
	defaultPort        = 8080
	title              = "Simple HTTP2 Server"
	defaultSSLKeyFile  = "/etc/ssl/private/lausanne_ch_2019.key"
	defaultSSLCertFile = "/etc/ssl/certs/lausanne_ch_2019.crt"
	htmlTemplate       = `
<!DOCTYPE html>
<html lang="en">
<head><meta name="viewport" content="width=device-width, initial-scale=1">
<meta charset="UTF-8">
<title>{{.TitlePage}}</title>
</head>
<body>
<h3>Headers</h3>
<p><ul>
{{range .Headers}}<li>{{ . }}</li>{{else}}<li><strong>no headers</strong></li>{{end}}
</ul></p>
</body>
</html>
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/skeleton/2.0.4/skeleton.min.css">
`
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// example : https://golang.org/src/net/http/example_test.go
	basePath := goutils.GetEnvOrDefault("GOPATH", "")
	SSLCertKeyFile := fmt.Sprintf("%s/%s", basePath, defaultSSLKeyFile)
	SSLCertificate := fmt.Sprintf("%s/%s", basePath, defaultSSLCertFile)
	port := defaultPort
	val, exist := os.LookupEnv("WEB_PORT")
	if !exist {
		flag.IntVar(&port, "port", defaultPort, "port the server will listen to")
	} else {
		golog.Info("Using ENV variable WEB_PORT to listen %s ", val)
		port, _ = strconv.Atoi(val)
	}
	listenAddr := fmt.Sprintf("%s:%v", defaultHost, port)

	golog.Info("\nlistening on  : %v\n", listenAddr)
	t := template.Must(template.New("page").Parse(htmlTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		golog.Info("connection from %s", r.RemoteAddr)
		var myStrings []string
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
		//Iterate over all header fields
		for k, v := range r.Header {
			tmp := fmt.Sprintf("Header field :  %q, Value %q\n", k, v)
			myStrings = append(myStrings, tmp)
		}
		/*
			fmt.Printf("myStrings : %T %v\n", myStrings, myStrings)
			fmt.Fprintf(w, "Host = %q\n", r.Host)
			fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
		*/
		//Get value for a specified token
		//fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])

		//fmt.Fprintf(w, "\n\nHello you, %q", html.EscapeString(r.URL.Path))

		htmlContent := htmlData{
			TitlePage: title,
			Headers:   myStrings,
		}
		err := t.Execute(w, htmlContent)
		check(err)
	})
	log.Fatal(http.ListenAndServeTLS(listenAddr, SSLCertificate, SSLCertKeyFile, nil))
}
