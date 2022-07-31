package main

import (
	"fmt"
	"net/http"
	"rsc.io/quote"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world!!, %s!", request.URL.Path[1:])
	fmt.Fprintf(writer, "%s", quote.Go())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
