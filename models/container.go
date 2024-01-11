package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
	"virtui/api"
)

var containersList []Container

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

type askCreateContainer struct {
	Name   string `json:"name"`
	Source struct {
		Type        string `json:"type"`
		Fingerprint string `json:"fingerprint"`
	} `json:"source"`
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

func GetContainersFromLocalStorage() ([]Container, error) {
	if containersList == nil {
		return nil, errors.New("local containers list is empty, before get from local storage try with api (GetContainersFromApi())")
	}
	return containersList, nil
}

func GetContainerWithName(name string) Container {
	GetContainersFromApi()
	return containersList[getIdContainerWithName(name)]
}

func CreateContainer(name string) Operation {
	var data askCreateContainer
	var operation Operation
	fingerprint := "1722a71a9f2dc0c68eac142a7d53ec728c15d2379e99f5b5545de99d440e3422"
	dataJson := fmt.Sprintf("{\"name\":\"%s\",\"source\":{\"type\":\"image\",\"fingerprint\":\"%s\"}}", name, fingerprint)
	err := json.Unmarshal([]byte(dataJson), &data)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(api.Cli.Post("/1.0/instances", data)), &operation)
	if err != nil {
		return Operation{}
	}
	return operation
}

// Copyright : NOAH MANDLER pour le nom de la fonction :3
func Exist(name string) bool {
	return getIdContainerWithName(name) != 0
}

func getIdContainerWithName(name string) int {
	for i, container := range containersList {
		if container.Metadata.Name == name {
			return i
		}
	}
	return 0
}

func DeleteContainerWithName(name string) (string, error) {
	GetContainersFromApi()
	if Exist(name) {
		return api.Cli.Delete(fmt.Sprintf("/1.0/instances/%s", name)), nil
	}
	return "", errors.New("Container doesn't exist")
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
	containersList = containersDetail
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
