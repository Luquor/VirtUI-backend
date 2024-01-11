package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"virtui/api"
)

type metaDataOperation struct {
	Id          string    `json:"id"`
	Class       string    `json:"class"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"`
	StatusCode  int       `json:"status_code"`
	Resources   struct {
		Containers []string `json:"containers"`
		Instances  []string `json:"instances"`
	} `json:"resources"`
	Metadata  interface{} `json:"metadata"`
	MayCancel bool        `json:"may_cancel"`
	Err       string      `json:"err"`
	Location  string      `json:"location"`
}

type Operation struct {
	Type       string            `json:"type"`
	Status     string            `json:"status"`
	StatusCode int               `json:"status_code"`
	Operation  string            `json:"operation"`
	ErrorCode  int               `json:"error_code"`
	Error      string            `json:"error"`
	Metadata   metaDataOperation `json:"metadata"`
}

type lastOperations struct {
	Type       string `json:"type"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Operation  string `json:"operation"`
	ErrorCode  int    `json:"error_code"`
	Error      string `json:"error"`
	Metadata   struct {
		Failure []string `json:"failure"`
		Running []string `json:"running"`
	} `json:"metadata"`
}

func OperationExist() bool {
	var operations lastOperations
	err := json.Unmarshal([]byte(api.Cli.Get("/1.0/operations")), &operations)
	if err != nil {
		log.Fatal(err)
	}
	return len(operations.Metadata.Running) > 0
}

func getOperationWithID(id string) Operation {
	var operationDetail Operation
	err := json.Unmarshal([]byte(api.Cli.Get(fmt.Sprintf("/1.0/operations/%s", id))), &operationDetail)
	if err != nil {
		log.Fatal(err)
	}
	return operationDetail
}

func GetLastOperation() Operation {
	var operations lastOperations
	var operationDetail Operation
	err := json.Unmarshal([]byte(api.Cli.Get("/1.0/operations")), &operations)
	if len(operations.Metadata.Running) == 0 {
		for _, metadatum := range operations.Metadata.Failure {
			err = json.Unmarshal([]byte(api.Cli.Get(metadatum)), &operationDetail)
			return operationDetail
		}
	}
	for _, metadatum := range operations.Metadata.Running {
		err = json.Unmarshal([]byte(api.Cli.Get(metadatum)), &operationDetail)
		return operationDetail
	}
	if err != nil {
		log.Fatal(err)
	}
	return Operation{}
}
