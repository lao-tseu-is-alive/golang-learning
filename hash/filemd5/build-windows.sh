#!/bin/sh
set -eu

script_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cd "$script_dir"

GOOS=windows GOARCH=386 go build -o filemd5.exe .
