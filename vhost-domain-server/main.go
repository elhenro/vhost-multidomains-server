package main

import (
	"fmt"
	"github.com/elhenro/testapp"
	"log"
	"net/http"
)

func helloA(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, testapp.Reverse("Hello, a.com"))
}

func helloSubA(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, sub.a.com")
}

func helloB(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, b.com")
}

func main() {
	http.HandleFunc("localhost/", helloA)
	http.HandleFunc("0.0.0.0/", helloSubA)
	http.HandleFunc("b.com/", helloB)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
