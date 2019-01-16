package main

import (
	"log"
	"net/http"
)

type banditResults struct {
	a []int     `json:a`
	b []float64 `json:b`
	c []float64 `json:c`
}

func main() {
	Do_bandit(2, []float64{0.1, 0.5}, 0.2, 10, 10)

	http.HandleFunc("/a", handler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.Handle("/", &templateHandler{filename: "resources/templates/index.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
