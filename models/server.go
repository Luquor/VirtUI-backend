package models

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"text/template"
	"virtui/api/modelsResponse"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	array := GetContainersFromApi()
	tmpl, err := template.ParseFiles("index.html")
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
	jsonResponse := modelsResponse.AddContainerResponse{}
	json.NewDecoder(r.Body).Decode(&jsonResponse)
	operation := CreateContainer(jsonResponse.Name, jsonResponse.Fingerprint)
	fmt.Println("Création d'un container (Status) : ", operation.Status, " ...")
	fmt.Println("Operation (status) :", GetOperationWithID(operation.Metadata.Id).Status)
	//fmt.Println("Create container", CreateContainer("Nouveau"+strconv.FormatInt(time.Now().Unix(), 10)))
	w.Write([]byte(fmt.Sprintf("Creating container... : %s", jsonResponse.Name)))
}

func getContainers(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting all the containers...")
	// w.Write(array)
	render.JSON(w, r, GetContainersFromApi())
}

func startContainer(w http.ResponseWriter, r *http.Request) {
	log.Print("Start a container...")
	name := chi.URLParam(r, "name")
	StartContainer(name)
	w.Write([]byte(fmt.Sprint("Starting... : ", name)))
}

func stopContainer(w http.ResponseWriter, r *http.Request) {
	log.Print("Stopping a container...")
	name := chi.URLParam(r, "name")
	StopContainer(name)
	w.Write([]byte(fmt.Sprint("Stopping... : ", name)))
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

func test(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, GetImages())
}

func StartWebServer() {

	log.Print("Starting web server...")

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Logger)

	// Image

	r.Get("/images", test)

	// Container

	r.Get("/", homepage)
	r.Post("/container", createContainer)
	r.Get("/containers", getContainers)
	r.Get("/container/{name}", getContainer)
	r.Put("/container/{name}/start", startContainer)
	r.Put("/container/{name}/stop", stopContainer)
	r.Delete("/container/{name}/", deleteContainer)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Server running on port 8000...")
}
