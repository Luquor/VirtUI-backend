package models

import (
	"encoding/json"
	"log"
	"time"
	"virtui/api"
)

type Container struct {
	Metadata struct {
		CreatedAt   time.Time `json:"created_at"`
		Description string    `json:"description"`
		Ephemeral   bool      `json:"ephemeral"`
		Location    string    `json:"location"`
		Name        string    `json:"name"`
		Restore     string    `json:"restore"`
		Status      string    `json:"status"`
		StatusCode  int       `json:"status_code"`
	} `json:"metadata"`
	api.StandardReturn
}

type containers struct {
	api.StandardReturn
	Operation string   `json:"operation"`
	ErrorCode int      `json:"error_code"`
	Error     string   `json:"error"`
	Metadata  []string `json:"metadata"`
}

// TO DO : Une vraie vérficiation svp
func IsContainerExist(nameC string) bool {
	return true
}

func GetContainersFromApi() []Container {
	var containersDetail []Container
	var containerDetail Container
	var containers containers
	err := json.Unmarshal([]byte(api.Cli.Get("/1.0/containers")), &containers)
	for _, metadatum := range containers.Metadata {
		err = json.Unmarshal([]byte(api.Cli.Get(metadatum)), &containerDetail)
		containersDetail = append(containersDetail, containerDetail)
	}
	if err != nil {
		log.Fatal(err)
	}
	return containersDetail
}

/**
func (c Containers) GetContainerByName(nameC string) *Container {
	var jsonDecode Container
	err := json.Unmarshal([]byte(api.Cli.Get(fmt.Sprintf("1.0/containers/%s", nameC))), &jsonDecode)
	if err != nil {
		log.Fatal(err)
	}
	return &jsonDecode
}
**/
