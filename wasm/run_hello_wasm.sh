#!/bin/bash
GOOS=js GOARCH=wasm go build -o hello.wasm hello_wasm.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
cp "$(go env GOROOT)/misc/wasm/wasm_exec.html" .
echo "navigate to : http://localhost:8080/hello_wasm.html"
go run miniweb.go
