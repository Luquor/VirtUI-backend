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
	"github.com/go-chi/cors"
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
	w.Write([]byte("Create container"))
}

func getContainers(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting all the containers...")
	// w.Write(array)
	render.JSON(w, r, GetContainersFromApi())
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

func getClusters(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting all the clusters...")
	// w.Write(array)
	render.JSON(w, r, GetClustersFromApi())
}

func getCluster(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting a cluster...")
	name := chi.URLParam(r, "serverName")
	dataJson, _ := GetClusterWithName(name)
	render.JSON(w, r, dataJson)
}

// func addClusterAddress(w http.ResponseWriter, r *http.Request) {
// 	cluster := chi.URLParam(r, "cluster")
// 	group := chi.URLParam(r, "group")
// 	log.Print("Creating a cluster...")

//		_, _ = AddClusterAddress(cluster, group)
//		w.Write([]byte("Create cluster"))
//	}
func createCluster(w http.ResponseWriter, r *http.Request) {
	log.Print("Creating a cluster...")
	group, _ := GetClusterGroupWithName(chi.URLParam(r, "group"))
	clusterName, _ := GetClusterWithName(chi.URLParam(r, "cluster"))

	_, _ = CreateCluster(group, clusterName)
	w.Write([]byte("Create cluster"))
}

func deleteCluster(w http.ResponseWriter, r *http.Request) {
	log.Print("Deleting a cluster...")
	name := chi.URLParam(r, "name")
	clusterName, _ := DeleteCluster(name)
	w.Write([]byte(clusterName))
}

func getContainerFromCluster(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting all the containers from a cluster...")
	name := chi.URLParam(r, "cluster")
	containerList, _ := GetContainersFromCluster(name)
	render.JSON(w, r, containerList)
}

func redirectToSpecificContainer(w http.ResponseWriter, r *http.Request) {
	log.Print("Redirecting to the container...")
	container := chi.URLParam(r, "container")
	http.Redirect(w, r, "/container/"+container, 301)
}

func StartWebServer() {
	fs := http.FileServer(http.Dir("static/stylesheets"))
	http.Handle("/static/stylesheets/", http.StripPrefix("/static/stylesheets/", fs))

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

	r.Get("/", homepage)
	r.Post("/container", createContainer)
	r.Get("/containers", getContainers)
	r.Get("/container/{name}", getContainer)
	r.Delete("/container/{name}/", deleteContainer)

	r.Get("/clusters", getClusters)
	r.Get("/cluster/{cluster}", getCluster)
	r.Post("/cluster", createCluster)
	r.Delete("/cluster/{cluster}", deleteCluster)
	r.Get("/cluster/{cluster}/container", getContainerFromCluster)
	r.Get("cluster/{cluster}/container/{container}", redirectToSpecificContainer)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Server running on port 8000...")
}
