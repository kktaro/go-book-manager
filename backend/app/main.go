package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":3000",
		Handler: nil,
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hello/test", requestGetter)

	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func requestGetter(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Name string `json:"name"`
	}
	type Result struct {
		Name string `json:"name"`
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Need Method: POST")
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Need ContentType: Json")
		return
	}

	request := &Request{}
	json.NewDecoder(r.Body).Decode(request)

	name := request.Name
	if name == "" {
		name = "Empty Name!"
	}

	result := Result{
		Name: name,
	}
	w.Header().Set("Content-Type", "application/json; charaset=UTF-8")
	json.NewEncoder(w).Encode(result)
}
