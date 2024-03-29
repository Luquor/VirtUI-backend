package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
	"virtui/api"
)

var snapshotList []Snapshot

type Snapshot struct {
	Metadata struct {
		Architecture string    `json:"architecture"`
		CreatedAt    time.Time `json:"created_at"`
		Ephemeral    bool      `json:"ephemeral"`
		Expires_at   time.Time `json:"expires_at"`
		Last_used_at time.Time `json:"last_used_at"`
		Name         string    `json:"name"`
		Profiles     []string  `json:"profiles"`
		Size         int       `json:"size"`
		Stateful     bool      `json:"stateful"`
	} `json:"metadata"`
	api.StandardReturn
}

// type snapshots struct {
// 	api.StandardReturn
// 	Operation string   `json:"operation"`
// 	ErrorCode int      `json:"error_code"`
// 	Error     string   `json:"error"`
// 	Metadata  []string `json:"metadata"`
// }

type askCreateSnapshot struct {
	Expires_at time.Time `json:"expires_at"`
	Name       string    `json:"name"`
	Stateful   bool      `json:"stateful"`
}

func CreateSnapshot(containerName string, snapshotName string) Operation {
	var data askCreateSnapshot
	var operation Operation
	if !IsContainerExist(containerName) {
		fmt.Println(errors.New("container does not exist"))
	}

	expirationDate := time.Now().AddDate(0, 0, 7)
	dataJson := fmt.Sprintf(`{"expires_at":"%s","name":"%s","stateful":true}`, expirationDate, snapshotName)

	err := json.Unmarshal([]byte(dataJson), &data)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(api.Cli.Post(fmt.Sprintf("/1.0/instances/%s/snapshots", containerName), data)), &operation)
	if err != nil {
		log.Fatal(err)
	}

	return operation
}

func RestoreSnapshot(containerName string, snapshot Snapshot) Operation {
	var operation Operation

	if !IsContainerExist(containerName) {
		fmt.Println(errors.New("container does not exist"))
	}
	if !IsSnapshotExist(snapshot.Metadata.Name) {
		fmt.Println(errors.New("snapshot does not exist"))
	}

	if snapshot.Metadata.Stateful {
		err := json.Unmarshal([]byte(api.Cli.Put(fmt.Sprintf("/1.0/instances/%s", containerName), fmt.Sprintf(`{"restore": "%s/%s", "stateful": true}`, containerName, snapshot.Metadata.Name))), &operation)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := json.Unmarshal([]byte(api.Cli.Put(fmt.Sprintf("/1.0/instances/%s", containerName), fmt.Sprintf(`{"restore": "%s/%s"}`, containerName, snapshot.Metadata.Name))), &operation)
		if err != nil {
			log.Fatal(err)
		}
	}

	return operation
}

func IsSnapshotExist(name string) bool {
	GetSnapshotsFromApi()
	for _, snapshot := range snapshotList {
		if snapshot.Metadata.Name == name {
			return true
		}
	}
	return false
}

func GetSnapshotsFromApi() {
	err := json.Unmarshal([]byte(api.Cli.Get("/1.0/instances")), &snapshotList)
	if err != nil {
		log.Fatal(err)
	}
}
