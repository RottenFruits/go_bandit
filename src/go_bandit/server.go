package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

type Result struct {
    N_arms     int  `json:"n_arms"`
    Arm_probs   struct {
        Prob      float64 `json:"prob"`
        Key int `json:"key"`
        Visible     bool `json:"visible"`
	} `json:"arm_probs"`
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

	
	//Validate request
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//parse json
	var jsonBody map[string]interface{}

	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Print(jsonBody)
	log.Print(jsonBody)
	
	//test(w, r)
}

