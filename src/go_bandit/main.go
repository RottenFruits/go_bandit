package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
	"fmt"
	"encoding/json"
)

type banditResults struct {
    a  []int `json:a`
    b  []float64 `json:b`
    c  []float64 `json:c`
}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(t.filename))
	})
	t.templ.Execute(w, nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	var arms BernoulliArms

	//log.Print(r)
	probs := [2]float64{0.4, 0.8}
	for _, p := range probs {
		arms = append(arms, BernoulliArm{p})
	}

	bandit := Bandit{}
	bandit.Initialize("EG", len(arms), 0.2)
	a, b, c := bandit.test_algorithm(arms, 10, 10)
	var res = banditResults{a, b, c}
	log.Print(res)
	//log.Print(b)
	/*
	res2, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	*/
	d, _ := json.Marshal(res)
	fmt.Fprintf(w, string(d))
	
	//w.Header().Set("Content-Type", "application/json")
	//log.Print(res)
	//fmt.Fprint(w, res2)
}

// リクエストを処理する関数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
	test(w, r)
}


func main() {
	//test()

	http.HandleFunc("/a", handler)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.Handle("/", &templateHandler{filename: "resources/templates/index.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
