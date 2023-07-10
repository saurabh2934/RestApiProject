package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello1) // http server register the function hello1
	http.ListenAndServe(":8081", nil)
}
func hello1(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello")
}
