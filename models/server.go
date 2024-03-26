package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"virtui/api/modelsResponse"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

const (
	STATUS_MOVED = http.StatusMovedPermanently
)

func homepage(w http.ResponseWriter, r *http.Request) {
	var array []Container
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
	array, err := GetContainersFromApi()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	render.JSON(w, r, array)
}

func getContainer(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting a container...")
	name := chi.URLParam(r, "name")

	container, err := GetContainerWithName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, container)
}

func consoleContainer(w http.ResponseWriter, r *http.Request) {
	log.Print("Get WebSocket Console")
	name := chi.URLParam(r, "name")
	test := GetConsoleForContainer(name)
	websocket := GetSocketsFromURLOperation(test.Operation)
	render.JSON(w, r, websocket)
}

func deleteContainer(w http.ResponseWriter, r *http.Request) {
	log.Print("Deleting a container...")
	name := chi.URLParam(r, "name")
	containerName, _ := DeleteContainerWithName(name)
	w.Write([]byte(containerName))
}

func getImages(w http.ResponseWriter, r *http.Request) {

	imagesList, err := GetImages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, imagesList)
}

func getClusters(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting all the clusters...")
	// w.Write(array)
	array, err := GetClustersFromApi()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	render.JSON(w, r, array)
}

func getCluster(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting a cluster...")
	name := chi.URLParam(r, "serverName")
	dataJson, err := GetClusterWithName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
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

func createContainerFromCluster(w http.ResponseWriter, r *http.Request) {
	log.Print("Creating a container from a cluster...")
	jsonResponse := modelsResponse.AddContainerResponse{}
	json.NewDecoder(r.Body).Decode(&jsonResponse)
	cluster := chi.URLParam(r, "cluster")
	_ = CreateContainerFromCluster(cluster, jsonResponse.Name, jsonResponse.Fingerprint)
	w.Write([]byte("Create container from cluster"))
}

func deleteContainerFromCluster(w http.ResponseWriter, r *http.Request) {
	log.Print("Deleting a container from a cluster...")
	cluster := chi.URLParam(r, "cluster")
	container := chi.URLParam(r, "container")
	_ = DeleteContainerFromCluster(cluster, container)
	w.Write([]byte("Delete container from cluster"))
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
	http.Redirect(w, r, "/container/"+container, STATUS_MOVED)
}

// container/{container}/actions (start, stop, restart => bodyparams)
func controlContainer(w http.ResponseWriter, r *http.Request) {
	log.Print("Control a container...")
	container := chi.URLParam(r, "container")
	r.ParseForm()
	jsonResponse := modelsResponse.ControlContainer{}
	json.NewDecoder(r.Body).Decode(&jsonResponse)
	response, err := ControlContainerWithName(container, jsonResponse.Action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(response))
}

func StartWebServer() {

	log.Print("Starting web server...")

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Image

	r.Get("/images", getImages)

	// Container

	r.Get("/", homepage)
	r.Post("/container", createContainer)
	r.Get("/containers", getContainers)
	r.Get("/container/{name}", getContainer)
	r.Delete("/container/{name}", deleteContainer)
	r.Get("/container/{name}/console", consoleContainer)

	r.Get("/clusters", getClusters)
	r.Get("/cluster/{cluster}", getCluster)
	r.Post("/cluster", createCluster)
	r.Delete("/cluster/{cluster}", deleteCluster)
	r.Post("/cluster/{cluster}/container", createContainerFromCluster)
	r.Delete("/cluster/{cluster}/container/{container}", deleteContainerFromCluster)
	r.Get("/cluster/{cluster}/container", getContainerFromCluster)
	r.Get("/cluster/{cluster}/container/{container}", redirectToSpecificContainer)
	r.Post("/container/{container}/actions", controlContainer)
	r.Post("/cluster/{cluster}/container/{container}/actions", controlContainer)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Server running on port 8000...")
}
