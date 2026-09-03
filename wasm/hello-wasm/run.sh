#!/bin/sh
set -eu

script_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cd "$script_dir"

GOOS=js GOARCH=wasm go build -o hello.wasm .
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
cp "$(go env GOROOT)/misc/wasm/wasm_exec.html" .
echo "navigate to: http://localhost:8080/index.html"
go run ../miniweb
