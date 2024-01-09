package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("static/stylesheets"))
		http.Handle("/static/stylesheets/", http.StripPrefix("/static/stylesheets/", fs))

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":80", nil)
}
