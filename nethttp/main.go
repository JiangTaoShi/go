package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":9080", nil)
}

func sayHello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello")
}
