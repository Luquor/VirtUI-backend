package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
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
	w.Write([]byte("Create container"))
}

func getContainers(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting all the containers...")
	array, _ := json.Marshal(GetContainersFromApi())
	w.Write(array)
}

func getContainer(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting a container...")
	name := chi.URLParam(r, "name")
	container, _ := json.Marshal(GetContainerWithName(name))
	w.Write(container)
}

func deleteContainer(w http.ResponseWriter, r *http.Request) {
	log.Print("Deleting a container...")
	name := chi.URLParam(r, "name")
	containerName, _ := DeleteContainerWithName(name)
	w.Write([]byte(containerName))
}

func StartWebServer() {
	fs := http.FileServer(http.Dir("static/stylesheets"))
	http.Handle("/static/stylesheets/", http.StripPrefix("/static/stylesheets/", fs))

	log.Print("Starting web server...")

	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Logger)

	r.Get("/", homepage)
	r.Post("/container", createContainer)
	r.Get("/containers", getContainers)
	r.Get("/container/{name}", getContainer)
	r.Delete("/container/{name}/", deleteContainer)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Server running on port 8000...")
}
