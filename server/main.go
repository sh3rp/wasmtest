package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "code.wasm")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must supply the base directory.")
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(os.Args[1])))
	mux.HandleFunc("/code.wasm", wasmHandler)
	log.Fatal(http.ListenAndServe(":8888", mux))
}
