package models

import (
	"encoding/json"
	"errors"
	"log"
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

func CreateCluster(serverName string, clusterAddress string, clusterPort int, clusterName string) {

}

func DeleteCluster(serverName string) (string, error) {
	getClustersFromApi()
	if clustersExist(serverName) {
		return api.Cli.Delete("/1.0/cluster/" + serverName), nil
	}
	return "", errors.New("Cluster does not exist")
}

func getClustersFromApi() []Cluster {
	var clustersDetail []Cluster
	var clusterDetail Cluster
	var clusters clusters
	err := json.Unmarshal([]byte(api.Cli.Get("/1.0/cluster")), &clusters)
	for _, metadatum := range clusters.Metadata {
		err = json.Unmarshal([]byte(api.Cli.Get(metadatum)), &clusterDetail)
		clustersDetail = append(clustersDetail, clusterDetail)
	}
	if err != nil {
		log.Fatal(err)
	}
	clustersList = clustersList
	return clustersDetail
}

func clustersExist(serverName string) bool {
	getClustersFromApi()
	for _, cluster := range clustersList {
		if cluster.Metadata.ServerName == serverName {
			return true
		}
	}
	return false
}
