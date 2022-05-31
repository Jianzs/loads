#!/bin/bash

BIN_DIR=${1-'.'}

go build -o ${BIN_DIR}/fib fib.go

# PAIRS=("linux-386" "linux-amd64" "linux-arm" "darwin-amd64" "windows-386" \
#        "windows-386" "windows-arm")
# for pair in ${PAIRS[@]}
# do
#     parts=(${pair//-/ })
#     os=${parts[0]}
#     arch=${parts[1]}
#     GOOS=$os GOARCH=$arch go build -o ${BIN_DIR}/fib-$os-$arch fib.go
# done