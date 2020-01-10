// [START gae_go111_app]

// cgil-golang-web-status is an App Engine GOLANG app.
package main

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	defaultHost        = ""
	defaultPort        = 8080
	defaultTitle       = "Simple Golang HTTP Server"
	defaultDescription = "Basic Golang HTTP Server"
	webRootDir         = "."
	htmlTemplate       = "resources/index_gotemplate.html"
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
	Path            string
	HttpProto       string
	HostRemote      string
	Host            string
	QueryParams     []keyValue
}

var webPort int

func check(msg string, err error) {
	if err != nil {
		log.Fatalf("ERROR : %s, error is [%v]", msg, err)
	}
}

func ServeStaticFileHandler(res http.ResponseWriter, req *http.Request) {
	// examine bundling res with app https://github.com/GeertJohan/go.rice
	fileName := path.Base(req.URL.Path)
	http.ServeFile(res, req, "./resources/"+fileName)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(webRootDir)))

	// beginning of parameters handling
	val, exist := os.LookupEnv("PORT")
	if !exist {
		webPort = defaultPort
	} else {
		log.Printf("Using ENV variable PORT : [%v] to listen \n", val)
		webPort, _ = strconv.Atoi(val)
	}

	mux := http.NewServeMux()
	listenAddr := fmt.Sprintf("%s:%v", defaultHost, webPort)
	log.Printf("#HTTP SERVER STARTED : listening on  %s\n", listenAddr)
	log.Println("#Method\tUrl\tProto\tPath\tRemoteAdr")

	//t := template.Must(template.New("page").ParseFiles(htmlTemplate))
	t, err := template.ParseFiles(htmlTemplate)
	check(fmt.Sprintf("FAILED TO PARSE %s", htmlTemplate), err)

	// ##### ROUTES #####
	// TODO:  handle 404 & bad host like ip
	mux.HandleFunc("/favicon.ico", ServeStaticFileHandler)
	mux.HandleFunc("/favicon-32x32.png", ServeStaticFileHandler)
	mux.HandleFunc("/robots.txt", ServeStaticFileHandler)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

		urlString, err := url.QueryUnescape(r.URL.String())
		if err != nil {
			log.Printf("ERROR : No way to decode url err is [%v] \n", err)
			urlString = html.UnescapeString("#URL_DECODE_ERROR#")
		}
		log.Printf("%s\t%s\t%s\t%s\t%s\n", r.Method, urlString, r.Proto, r.URL.Path, r.RemoteAddr)

		htmlContent := htmlData{
			TitlePage:       title,
			DescriptionPage: description,
			Headers:         myStrings,
			HttpMethod:      r.Method,
			HttpURL:         urlString,
			Path:            r.URL.Path,
			HttpProto:       r.Proto,
			Host:            r.Host,
			HostRemote:      r.RemoteAddr,
			QueryParams:     queryParams,
		}
		err = t.Execute(w, htmlContent)
		check("Template execute failure", err)
	})
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	srv := &http.Server{
		Addr:           listenAddr,
		Handler:        mux,
		TLSNextProto:   nil, // to allow http2 to run :->https://golang.org/pkg/net/http/
		ErrorLog:       logger,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 << 20 specifies maximum of 1MB header .
	}
	err = srv.ListenAndServe()
	check("Server ListenAndServe failure ", err)
}
