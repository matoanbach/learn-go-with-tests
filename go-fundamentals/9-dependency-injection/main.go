package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(os.Stdout, "Hello, %s\n", name)
}
func MyGreeterHandler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Greet(w, name)
	}
}

func main() {
	name := "World"
	args := os.Args
	if len(args) > 1 {
		name = args[1]
	}
	log.Fatal(http.ListenAndServe(":5001", MyGreeterHandler(name)))
}
