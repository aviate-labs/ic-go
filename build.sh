#!/usr/bin/env sh

GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o src/example/assets/main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" src/example/src/wasm_exec.js
