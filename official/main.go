package main

import (
	"net/http"
	"fmt"
)

func send(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", send)
	s := http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	fmt.Println(s.ListenAndServe())
}
