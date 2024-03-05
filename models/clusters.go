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

type ClusterGroup struct {
	Members []string `json:"members"`
	Name    string   `json:"name"`
}

type clusters struct {
	api.StandardReturn
	Operation string   `json:"operation"`
	ErrorCode int      `json:"error_code"`
	Error     string   `json:"error"`
	Metadata  []string `json:"metadata"`
}

func AddClusterAdress(cluster Cluster, group ClusterGroup) (string, error) {
	getClustersFromApi()
	if clustersExist(cluster.Metadata.ServerName) {
		return "", errors.New("Cluster already exists")
	}
	groups := group
	groups.Members = append(groups.Members, cluster.Metadata.ServerName)
	return api.Cli.Post("/1.0/cluster/groups", groups), nil
}

func DeleteCluster(serverName string) (string, error) {
	getClustersFromApi()
	if clustersExist(serverName) {
		return api.Cli.Delete("/1.0/cluster/members/" + serverName), nil
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
	clustersList = clustersDetail
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
