all: build run

run:
	go run server/main.go .

build:
	GOOS=js GOARCH=wasm go build -o code.wasm js/main.go

.PHONY: run buildasm 
