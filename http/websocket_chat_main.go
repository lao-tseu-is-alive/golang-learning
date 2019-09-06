// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "./html/websockets_home.html")
}

/*
	this example comes from https://github.com/gorilla/websocket/tree/master/examples/chat
	it shows how to use the websocket package to implement a simple web chat application
	using Gorilla's Go implementation of the WebSocket protocol : https://github.com/gorilla/websocket
	to run :
	cd ${GOPATH}/src/github.com/lao-tseu-is-alive/golang-learning/http
	go get -u github.com/gorilla/websocket
	go run websocket_chat_*.go
*/

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
