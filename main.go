package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"virtui/models"
)

func main() {
	//test := models.CreateContainer("JeSuisLUCAS")
	//fmt.Println(test)
	webServer()
}

func webServer() {
	fs := http.FileServer(http.Dir("static/stylesheets"))
	http.Handle("/static/stylesheets/", http.StripPrefix("/static/stylesheets/", fs))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		array := models.GetContainersFromApi()
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, array)
		if err != nil {
			log.Fatal(err)
		}
	})

	r.Post("/container", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("Create container", models.CreateContainer("Nouveau"+strconv.FormatInt(time.Now().Unix(), 10)))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	})

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
