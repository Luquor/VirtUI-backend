package models

import (
	"time"
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

type lastOperation struct {
	Type       string `json:"type"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Operation  string `json:"operation"`
	ErrorCode  int    `json:"error_code"`
	Error      string `json:"error"`
	Metadata   struct {
		Failure []metaDataOperation `json:"failure"`
		Running []metaDataOperation `json:"running"`
	} `json:"metadata"`
}

func GetLastOperation() {
}
