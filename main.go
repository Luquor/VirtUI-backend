package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"virtui/models"
)

type Container struct {
	Header string
	Status string
}

func main() {
	for _, value := range models.GetContainersFromApi() {
		fmt.Printf("%s\n", value.Metadata.Name)
	}

	fmt.Print(models.GetContainerWithName("noah"))

	//webServer()

}

func webServer() {
	fs := http.FileServer(http.Dir("static/stylesheets"))
	http.Handle("/static/stylesheets/", http.StripPrefix("/static/stylesheets/", fs))

	array := []Container{
		Container{Header: "Portainer", Status: "Stoppé"},
		Container{Header: "Grafana", Status: "Démarré"},
		Container{Header: "Nginx Proxy Manager", Status: "Stoppé"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}

		if r.Method == "POST" && r.URL.Path == "/" {
			r.ParseForm()
			header := "Nouveaux containers"
			status := "Démarré"

			container := Container{Header: header, Status: status}
			array = append(array, container)

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		err = tmpl.Execute(w, array)
		if err != nil {
			log.Fatal(err)
		}
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
