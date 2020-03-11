package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Server Running...")
	http.ListenAndServe(":3000", nil)
}
