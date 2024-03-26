package models

import (
	"encoding/json"
	"time"
	"virtui/api"
)

type Console struct {
	Metadata struct {
		Class       string    `json:"class"`
		CreatedAt   time.Time `json:"created_at"`
		Description string    `json:"description"`
		Err         string    `json:"err"`
		Id          string    `json:"id"`
		Location    string    `json:"location"`
		MayCancel   bool      `json:"may_cancel"`
		Metadata    struct {
			Command     []string    `json:"command"`
			Environment interface{} `json:"environment"`
			Fds         interface{}
			Interactive bool `json:"interactive"`
		} `json:"metadata"`
		Resources struct {
			Containers []string `json:"containers"`
			Instances  []string `json:"instances"`
		} `json:"resources"`
		Status     string    `json:"status"`
		StatusCode int       `json:"status_code"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"metadata"`
	Operation string `json:"operation"`
	api.StandardReturn
}

type consoleType struct {
	Command     []string `json:"command"`
	Cwd         string   `json:"cwd"`
	Environment struct {
		TERM string `json:"TERM"`
		HOME string `json:"HOME"`
	} `json:"environment"`
	Interactive      bool `json:"interactive"`
	Group            int  `json:"group"`
	User             int  `json:"user"`
	WaitForWebsocket bool `json:"wait-for-websocket"`
}

func GetConsoleForContainer(containerName string) (Console, error) {
	var console Console
	var data consoleType
	dataJson := "{\"command\": [\"bash\"],\"cwd\": \"/root\",\"environment\": {\"TERM\": \"xterm-256color\",\"HOME\": \"/root\"},\"interactive\": true,\"group\": 0,\"user\": 0,\"wait-for-websocket\": true}"
	err := json.Unmarshal([]byte(dataJson), &data)
	err = json.Unmarshal([]byte(api.Cli.Post("/1.0/instances/"+containerName+"/exec", data)), &console)
	return console, err

}
