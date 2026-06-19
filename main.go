package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server is listening")
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/hello", AliceHandler)
	http.HandleFunc("/count", CountHandler)
	http.ListenAndServe("localhost:8080", nil)
}
