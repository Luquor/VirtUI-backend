package models

import (
	"virtui/api"
)

var clustersList []Cluster

type Cluster struct {
	Metadata struct {
		Enabled      bool `json:"enabled"`
		MemberConfig []struct {
			ClusterAddress string `json:"cluster_address"`
			ClusterPort    int    `json:"cluster_port"`
			ClusterName    string `json:"cluster_name"`
		} `json:"member_config"`
		ServerName string `json:"server_name"`
	} `json:"metadata"`
	api.StandardReturn
}

type clusters struct {
	api.StandardReturn
	Operation string   `json:"operation"`
	ErrorCode int      `json:"error_code"`
	Error     string   `json:"error"`
	Metadata  []string `json:"metadata"`
}
