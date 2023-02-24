package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lao-tseu-is-alive/golog"
	"github.com/lao-tseu-is-alive/goutils"
	"html"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"strconv"
	"strings"
	"time"
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

type Book struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Isbn10  string `json:"isbn"`
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
.logo{ float: left;padding-right: 1em;}
</style>
</head>
<body>
<div class="container">
<img class="logo" src="/logo.svg"/>
<h3>{{.TitlePage}}</h3>
<h4>Request Info</h4>
<div class="row">
    <div class="three columns"><strong>Proto:</strong> {{.HttpProto}}</div>
	<div class="three columns"><strong>Method:</strong> {{.HttpMethod}}</div>
    <div class="three columns"><strong>Host:</strong> {{.Host}}</div>
	<div class="three columns"><strong>Remote IP:</strong> {{.HostRemote}}</div>
</div>
<div class="row">
    <div class="three columns"><strong>Path:</strong> {{.Path}}</div>	
    <div class="nine columns"><strong>Url:</strong> {{.HttpURL}}</div>	
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

func ServeStaticFileHandler(res http.ResponseWriter, req *http.Request) {
	// examine bundling res with app https://github.com/GeertJohan/go.rice
	fileName := path.Base(req.URL.Path)
	fullPath := "./resources/" + fileName
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		golog.Err("resource does not exist : %s ", fullPath)
	} else {
		golog.Info("serving %v", fullPath)
		http.ServeFile(res, req, fullPath)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var (
		myStrings   []keyValue
		queryParams []keyValue
		title       string = defaultTitle
		description string = defaultDescription
	)
	t := template.Must(template.New("page").Parse(htmlTemplate))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//Get value for a specified token
	//fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])

	//Iterate over all header fields
	for k, v := range r.Header {
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

	htmlContent := htmlData{
		TitlePage:       title,
		DescriptionPage: description,
		Headers:         myStrings,
		HttpMethod:      r.Method,
		Path:            r.URL.Path,
		HttpProto:       r.Proto,
		Host:            r.Host,
		HostRemote:      r.RemoteAddr,
		QueryParams:     queryParams,
	}
	err := t.Execute(w, htmlContent)
	check(err)
}

/*
*
will handle the creation of a "new" Book entry
you can test with :
curl -kvv -H "Content-Type: application/json" -d '{"title":"hello world", "authors":"Harold Wellington"}' https://localhost:8080/books/21

	*
*/
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	var data Book
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	title := strings.TrimSpace(data.Title)
	// title should be at least 3 chars long
	if len(title) < 3 {
		msg := fmt.Sprintf("invalid title parameter")
		golog.Err(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	authors := data.Authors
	isbn10 := data.Isbn10
	aBook := Book{
		Id:      0,
		Title:   title,
		Authors: authors,
		Isbn10:  isbn10,
	}
	js, err := json.Marshal(aBook)
	if err != nil {
		golog.Err("problem doing json.Marshal")
		http.Error(w, "ERROR CreateBook : "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
func ReadBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 0 {
		msg := fmt.Sprintf("invalid books id parameter  err: %v", err)
		golog.Err(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	golog.Info("INSIDE ReadBook(%v)", id)
	aBook := Book{
		Id:      id,
		Title:   "The \"Go\" Programming Language",          // note the double quote char => invalid json
		Authors: "Alan A. A. Donovan,\n Brian W. Kernighan", // same with newline
		Isbn10:  "9780134190440",
	}
	data, err := json.Marshal(aBook)
	if err != nil {
		msg := fmt.Sprintf("problem doing json.Marshal error : %v", err)
		golog.Err(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	golog.Info("about to return json : %v", string(data))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

/*
very basic middleware which logs the URI of the request being handled could be written as:
*/
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		urlString, err := url.QueryUnescape(r.URL.String())
		if err != nil {
			golog.Err("No way to decode url err: %v ", err)
			urlString = html.UnescapeString("#URL_DECODE_ERROR#")
		}
		logLine := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\n", r.Method, urlString, r.Proto, r.URL.Path, r.RemoteAddr, r.Host)
		golog.Info(logLine)
		log.Println(fmt.Sprintf("[%s]\t%s", golog.GetTimeStamp(), logLine))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
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
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*5, "the duration for which the server gracefully wait for existing connections to finish - e.g. 5s or 1m")
	flag.Parse()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	listenAddr := fmt.Sprintf("%s:%v", defaultHost, port)
	listenInfo := fmt.Sprintf("### Server started and listening on  : %v\n", listenAddr)
	logger.Printf(listenInfo)
	golog.Info(listenInfo)
	gomux := mux.NewRouter()
	logger.Println("#Method\tUrl\tProto\tPath\tRemoteAdr")

	// ##### ROUTES #####
	// TODO:  handle 404 & bad host like ip
	gomux.HandleFunc("/favicon.ico", ServeStaticFileHandler)
	gomux.HandleFunc("/logo.svg", ServeStaticFileHandler)
	gomux.HandleFunc("/favicon-32x32.png", ServeStaticFileHandler)
	gomux.HandleFunc("/robots.txt", ServeStaticFileHandler)
	gomux.HandleFunc("/books/{id}", CreateBook).Methods("POST")
	gomux.HandleFunc("/books/{id}", ReadBook).Methods("GET")
	/*
		mux.HandleFunc("/books/{id}",  UpdateBook).Methods("PUT")
		mux.HandleFunc("/books/{id}",  DeleteBook).Methods("DELETE")
	*/
	gomux.HandleFunc("/", HomeHandler)
	gomux.Use(loggingMiddleware)

	// ## let's run the server
	srv := &http.Server{
		Addr:           listenAddr,
		Handler:        gomux,
		TLSNextProto:   nil, // to allow http2 to run :->https://golang.org/pkg/net/http/
		ErrorLog:       logger,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 << 20 specifies maximum of 1MB header .
	}
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		err := srv.ListenAndServeTLS(SSLCertificate, SSLCertKeyFile)
		if err != nil {
			msg := fmt.Sprintf("ERROR when ListenAndServeTLS : %v", err)
			golog.Err(msg)
			log.Println(msg)
			os.Exit(1)
		}
	}()
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("### shutting down the http server...")
	os.Exit(0)
}
