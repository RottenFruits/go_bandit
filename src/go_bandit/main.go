package main

import (
	"log"
	"net/http"
)

func main() {
	//Do_bandit(2, []float64{0.1, 0.5}, 0.2, 10, 10)

	http.HandleFunc("/a", handler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.Handle("/", &templateHandler{filename: "resources/templates/index.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
