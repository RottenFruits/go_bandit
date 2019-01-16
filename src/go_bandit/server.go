package main

import (
	"encoding/json"
	//"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

type allParameter struct {
	Bandit        []Bandit        `json:"bandit"`
	ArmPrameters  []armPrameters  `json:"arm_parameters"`
	BanditResults []banditResults `json:"bandit_results"`
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

// リクエストを処理する関数
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello World from Go.")

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
	//var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &all_parameter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//fmt.Fprint(w, all_parameter)
	var probs []float64
	for _, para := range all_parameter.ArmPrameters {
		probs = append(probs, para.Prob)
	}

	//Do_bandit(all_parameter.Bandit[0].N, probs, 0.5, 1, 500)
	all_parameter.BanditResults[0] = Oneshot_bandit(&all_parameter.Bandit[0], all_parameter.BanditResults[0], probs, 0.2)
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(all_parameter)
	w.Write(res)

	//fmt.Fprint(w, res)
}
