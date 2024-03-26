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

// Action possible : start, stop, restart, freeze, unfreeze
// type askState struct {
// 	Action   string `json:"action"`
// 	Force    bool   `json:"force"`
// 	Stateful bool   `json:"stateful"`
// 	Timeout  int    `json:"timeout"`
// }

type containers struct {
	api.StandardReturn
	Operation string   `json:"operation"`
	ErrorCode int      `json:"error_code"`
	Error     string   `json:"error"`
	Metadata  []string `json:"metadata"`
}

func IsContainerExist(name string) bool {
	GetContainersFromApi()
	for _, container := range containersList {
		if container.Metadata.Name == name {
			return true
		}
	}
	return false
}

func GetContainersFromLocalStorage() ([]Container, error) {
	if containersList == nil {
		return nil, errors.New("local containers list is empty, before get from local storage try with api (GetContainersFromApi())")
	}
	return containersList, nil
}

func GetContainerWithName(name string) (Container, error) {
	GetContainersFromApi()
	if len(containersList) == 0 {
		return Container{}, errors.New("the container list is empty")
	}
	if !IsContainerExist(name) {
		return Container{}, errors.New("container doesn't exist")
	}

	container := containersList[getIdContainerWithName(name)]
	return container, nil
}

func CreateContainer(name string, fingerprint string) Operation {
	var data askCreateContainer
	var operation Operation
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

func exist(name string) bool {
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
	if exist(name) {
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

func StartContainer(name string) (string, error) {
	if !IsContainerExist(name) {
		log.Fatal("container doesn't exist")
	}
	container, _ := GetContainerWithName(name)

	if container.Metadata.Status == "Running" {
		return "", errors.New("container is already running")
	}
	return api.Cli.Put(fmt.Sprintf("/1.0/instances/%s/state", name), "{\"action\":\"start\"}"), nil
}

func StopContainer(name string) (string, error) {
	if !IsContainerExist(name) {
		log.Fatal("Container doesn't exist")
	}
	container, _ := GetContainerWithName(name)

	if container.Metadata.Status == "Stopped" {
		return "", errors.New("container is already stopped")
	}
	return api.Cli.Put(fmt.Sprintf("/1.0/instances/%s/state", name), "{\"action\":\"stop\"}"), nil
}

func RestartContainer(name string) (string, error) {
	if !IsContainerExist(name) {
		log.Fatal("Container doesn't exist")
	}
	container, _ := GetContainerWithName(name)
	if container.Metadata.Status == "Stopped" {
		return "", errors.New("container is already stopped")
	}
	return api.Cli.Post(fmt.Sprintf("/1.0/instances/%s/state", name), "{\"action\":\"restart\"}"), nil
}

func ControlContainerWithName(name string, action string) (string, error) {
	switch action {
	case "start":
		return StartContainer(name)
	case "stop":
		return StopContainer(name)
	case "restart":
		return RestartContainer(name)
	default:
		return "", errors.New("action not found")
	}
}
