#!/usr/bin/env sh

GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o src/go_assets/assets/main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" src/go_assets/src/wasm_exec.js
