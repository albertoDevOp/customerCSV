package main

import {
	"fmt",
	"net/http"
}

func main() {

	http.HandleFunc('/', func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}