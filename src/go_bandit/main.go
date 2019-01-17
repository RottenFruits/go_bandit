package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/oneshot", Handler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.Handle("/", &templateHandler{filename: "resources/templates/index.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
