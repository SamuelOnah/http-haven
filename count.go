package main

import (
	"fmt"
	"io"
	"net/http"
)

func CountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, "Send a POST request with text to count words")
	}
	if r.Method == "POST" {
		count, _ := io.ReadAll((r.Body))
		// if err != nil {
		// 	http.Error(w, )
		// }
		re := len(string(count))
		fmt.Fprint(w, re)
	}

}
