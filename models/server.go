package models

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	array := GetContainersFromApi()
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, array)
	if err != nil {
		log.Fatal(err)
	}
}

func createContainer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Create container", CreateContainer("Nouveau"+strconv.FormatInt(time.Now().Unix(), 10)))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func StartWebServer() {
	fs := http.FileServer(http.Dir("static/stylesheets"))
	http.Handle("/static/stylesheets/", http.StripPrefix("/static/stylesheets/", fs))

	log.Print("Starting web server...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", homepage)
	r.Post("/container", createContainer)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Server running on port 8000...")
}
