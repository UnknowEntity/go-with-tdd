package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(write io.Writer, name string) {
	fmt.Fprintf(write, "Hello %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
