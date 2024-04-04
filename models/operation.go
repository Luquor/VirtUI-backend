package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
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

type WebSocketConsole struct {
	Data     string
	Control  string
	Location string
}

func OperationExist() bool {
	var operations lastOperations
	err := json.Unmarshal([]byte(api.Cli.Get("/1.0/operations")), &operations)
	if err != nil {
		log.Fatal(err)
	}
	return len(operations.Metadata.Running) > 0
}

func GetOperationWithID(id string) Operation {
	var operationDetail Operation
	err := json.Unmarshal([]byte(api.Cli.Get(fmt.Sprintf("/1.0/operations/%s", id))), &operationDetail)
	if err != nil {
		log.Fatal(err)
	}
	return operationDetail
}

func GetOperationWithURL(url string) (Operation, error) {
	var operationDetail Operation
	err := json.Unmarshal([]byte(api.Cli.Get(url)), &operationDetail)

	return operationDetail, err
}

func GetSocketsFromURLOperation(url string) (WebSocketConsole, error) {
	operation, err := GetOperationWithURL(url)
	parts := strings.Split(url, "?")
	url = parts[0]
	token := operation.Metadata.Metadata.(map[string]interface{})["fds"].(map[string]interface{})
	return WebSocketConsole{Data: fmt.Sprintf("%s/websocket?secret=%s", url, token["0"]), Control: fmt.Sprintf("%s/websocket?secret=%s", url, token["control"])}, err
}

func GetLastOperation() Operation {
	var operations lastOperations
	var operationDetail Operation
	err := json.Unmarshal([]byte(api.Cli.Get("/1.0/operations")), &operations)
	if len(operations.Metadata.Running) == 0 {
		for _, metadatum := range operations.Metadata.Failure {
			err = json.Unmarshal([]byte(api.Cli.Get(metadatum)), &operationDetail)
			if err != nil {
				log.Fatal(err)
			}
			return operationDetail
		}
	}
	for _, metadatum := range operations.Metadata.Running {
		err = json.Unmarshal([]byte(api.Cli.Get(metadatum)), &operationDetail)
		if err != nil {
			log.Fatal(err)
		}
		return operationDetail
	}
	if err != nil {
		log.Fatal(err)
	}
	return Operation{}
}
