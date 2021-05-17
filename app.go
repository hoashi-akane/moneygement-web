package main

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func main() {
	http.HandleFunc("/", http.HandlerFunc(handler))
	http.ListenAndServe(":8080", nil)
}
