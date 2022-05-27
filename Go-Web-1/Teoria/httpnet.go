package main

// Criando um servidor em GO

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello World!`))
	// outra forma de dar o print fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// endpoint - Local onde vamos acessar a função.
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
