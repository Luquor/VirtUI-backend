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

// Il faut créer la fonction CreateCluster !
// Trouver comment créer un cluster avec l'api !
// // Il n'y a que que un de quoi créer un cluster group et non juste un seul.
// func CreateCluster(serverName string, clusterAddress string, clusterPort int, clusterName string) {
// 	 cluster := Cluster{}
// 	 dataJson := `{"metadata": {"server_name": "` + serverName + `", "enabled": true, "member_config": [{"cluster_address": "` + clusterAddress + `", "cluster_port": ` + string(clusterPort) + `, "cluster_name": "` + clusterName + `"}]}}`
// 	 err := json.Unmarshal([]byte(api.Cli.Post("/1.0/cluster", dataJson)), &cluster)
// 	 if err != nil {
// 	 	log.Fatal(err)
// 	 }

// }

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
