package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"sync"
	//"log"
	//"fmt"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

type allParameter struct {
	Bandit       Bandit         `json:"bandit"`
	ArmPrameters []armPrameters `json:"arm_parameters"`
}

type armPrameters struct {
	Prob    float64 `json:"prob"`
	Key     int     `json:"key"`
	Visible bool    `json:"visible"`
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(t.filename))
	})
	t.templ.Execute(w, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
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
	var all_parameter allParameter
	err = json.Unmarshal(body[:length], &all_parameter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//log.Print(all_parameter)
	var probs []float64
	for _, para := range all_parameter.ArmPrameters {
		probs = append(probs, para.Prob)
	}

	Oneshot_bandit(&all_parameter.Bandit, probs, 0.2)

	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(all_parameter)
	w.Write(res)

	//fmt.Fprint(w, res)
}
