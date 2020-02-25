package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	http.ListenAndServe(":8080", nil)
}

func handleHello(rw http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
}
