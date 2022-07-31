package main

import (
	"fmt"
	"log"
	"net/http"

	"go-sample/greetings"

	"rsc.io/quote"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world!!, %s!", request.URL.Path[1:])
	fmt.Fprintf(writer, "%s", quote.Go())
}

func write() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

  names := []string {"akira", "moroishi", "moroaki"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

func main() {
	write()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
