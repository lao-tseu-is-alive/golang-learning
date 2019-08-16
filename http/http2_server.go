package main

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"github.com/lao-tseu-is-alive/goutils"
	"html"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

type keyValue = struct {
	Key   string
	Value string
}

type htmlData = struct {
	TitlePage       string
	DescriptionPage string
	Headers         []keyValue
	HttpMethod      string
	HttpURL         string
	HttpProto       string
	HostRemote      string
	QueryParams     []keyValue
}

const (
	defaultHost        = "localhost"
	defaultPort        = 8080
	defaultTitle       = "Simple Golang HTTP2 Server"
	defaultDescription = "Basic Golang HTTP2 Server"
	defaultSSLKeyFile  = "/etc/ssl/private/lausanne_ch_2019.key"
	defaultSSLCertFile = "/etc/ssl/certs/lausanne_ch_2019.crt"
	htmlTemplate       = `
<!DOCTYPE html>
<html lang="en">
<head><meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
<meta charset="UTF-8">
<meta name="Description" content="{{.DescriptionPage}}">
<title>{{.TitlePage}}</title>
<style>
.bold{ font-weight: bold;}
</style>
</head>
<body>
<div class="container">
<h3>{{.TitlePage}}</h3>
<h4>Request Info</h4>
<div class="row">
    <div class="three columns"><strong>Proto:</strong> {{.HttpProto}}</div>
	<div class="three columns"><strong>Method:</strong> {{.HttpMethod}}</div>
    <div class="three columns"><strong>Url:</strong> {{.HttpURL}}</div>
	<div class="three columns"><strong>Remote IP:</strong> {{.HostRemote}}</div>
</div>
<h4>Headers</h4>
<p><ul>
{{range .Headers}}<li><span class='bold'>{{ .Key }} :</span> {{ .Value }}</li>{{else}}<li><strong>no headers</strong></li>{{end}}
</ul></p>
<h4>Query params</h4>
<p><ul>
{{range .QueryParams}}<li><span class='bold'>{{ .Key }} :</span> {{ .Value }}</li>{{else}}<li><strong>no params</strong></li>{{end}}
</ul></p>
</div>
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

/*
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "resources/favicon.ico")
}

func faviconPngHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "resources/favicon-32x32.png")
}
*/
func ServeStaticFileHandler(res http.ResponseWriter, req *http.Request) {
	fileName := path.Base(req.URL.Path)
	http.ServeFile(res, req, "./resources/"+fileName)
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
	mux := http.NewServeMux()
	golog.Info("listening on  : %v\n", listenAddr)
	fmt.Println("#Method\tUrl\tProto\tPath\tRemoteAdr")

	t := template.Must(template.New("page").Parse(htmlTemplate))
	// ##### ROUTES #####
	mux.HandleFunc("/favicon.ico", ServeStaticFileHandler)
	mux.HandleFunc("/favicon-32x32.png", ServeStaticFileHandler)
	mux.HandleFunc("/robots.txt", ServeStaticFileHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// golog.Info("connection from %s", r.RemoteAddr)
		var (
			myStrings   []keyValue
			queryParams []keyValue
			title       string = defaultTitle
			description string = defaultDescription
		)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//Get value for a specified token
		//fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])
		//Iterate over all header fields
		for k, v := range r.Header {
			// tmp := fmt.Sprintf("%q : %q\n", k, v)
			kv := keyValue{
				Key:   k,
				Value: strings.Join(v, ","),
			}
			myStrings = append(myStrings, kv)
		}
		//Iterate over all header fields
		for k, v := range r.URL.Query() {
			// tmp := fmt.Sprintf("%q : %v\n", k, v)
			kv := keyValue{
				Key:   k,
				Value: strings.Join(v, ","),
			}
			queryParams = append(queryParams, kv)
		}
		//fmt.Fprintf(w, "Host = %q\n", r.Host)

		urlString, err := url.QueryUnescape(r.URL.String())
		if err != nil {
			golog.Err("No way to decode url err: %v ", err)
			urlString = html.UnescapeString("#URL_DECODE_ERROR#")
		}
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n", r.Method, urlString, r.Proto, r.URL.Path, r.RemoteAddr)
		fmt.Println("URL.Query : ", r.URL.Query())

		htmlContent := htmlData{
			TitlePage:       title,
			DescriptionPage: description,
			Headers:         myStrings,
			HttpMethod:      r.Method,
			HttpURL:         urlString,
			HttpProto:       r.Proto,
			HostRemote:      r.RemoteAddr,
			QueryParams:     queryParams,
		}
		err = t.Execute(w, htmlContent)
		check(err)
	})
	srv := &http.Server{
		Addr:         listenAddr,
		Handler:      mux,
		TLSNextProto: nil, // to allow http2 to run !
	}

	log.Fatal(srv.ListenAndServeTLS(SSLCertificate, SSLCertKeyFile))
}
