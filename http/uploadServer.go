package main

import (
	"flag"
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const uploadHtmlFilePath = "html/basicupload.html"

func uploadFile(w http.ResponseWriter, r *http.Request) {
	golog.Un(golog.Trace("File Upload Endpoint Hit"))

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		golog.Err("Error Retrieving the uploaded File err : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	extension := filepath.Ext(handler.Filename)
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//TODO: check that extension corresponds to "real" file content

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp_upload", fmt.Sprintf("upload-*%s", extension))
	if err != nil {
		errMsg := fmt.Sprintf("Error creating temp file :%v", err)
		golog.Err(errMsg)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errMsg)
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		errMsg := fmt.Sprintf("Error reading uploaded file :%v", err)
		golog.Err(errMsg)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errMsg)
		return
	}
	// write this byte array to our temporary file
	n, err := tempFile.Write(fileBytes)
	if err != nil {
		errMsg := fmt.Sprintf("Error writing temp file :%v", err)
		golog.Err(errMsg)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errMsg)
		return
	}
	// return that we have successfully uploaded our file!
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Successfully Uploaded File (%d bytes written) in %s\n", n, tempFile.Name())
	return
}

func serveHtmlFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, uploadHtmlFilePath)
}

func setupRoutes(listenAdr string) {
	http.HandleFunc("/go_upload", uploadFile)
	http.HandleFunc("/", serveHtmlFile)
	log.Fatal(http.ListenAndServe(listenAdr, nil))
}

func main() {
	const defaultHost = "localhost"
	const defaultPort = 8080
	port := defaultPort
	val, exist := os.LookupEnv("WEB_PORT")
	if !exist {
		flag.IntVar(&port, "port", defaultPort, "port the server will listen to")
	} else {
		golog.Info("Using ENV variable WEB_PORT to listen %s ", val)
		port, _ = strconv.Atoi(val)
	}

	listenAddress := fmt.Sprintf("%s:%v", defaultHost, port)
	golog.Info("\nlistening on : %s\n", listenAddress)
	setupRoutes(listenAddress)
}
