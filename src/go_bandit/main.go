package main

import (
	"log"
	"net/http"
	"html/template"
	"sync"
	"path/filepath"

)

type templateHandler struct{
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.once.Do(func(){
		t.templ =
		template.Must(template.ParseFiles(filepath.Join("templates",
		t.filename)))
	})
	t.templ.Execute(w, nil)
}



func test(){
	var arms BernoulliArms
	probs := [2]float64{0.4, 0.8}
	for _, p := range probs {
		arms = append(arms, BernoulliArm{p})
	}

	bandit := Bandit{}
	bandit.Initialize("EG", len(arms), 0.2)
	a, b, c := bandit.test_algorithm(arms, 50, 500)
	log.Print(a)
	log.Print(b)
	log.Print(c)
}

func main() {
	//test()
	
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img/"))))
	http.Handle("/", &templateHandler{filename: "index.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}