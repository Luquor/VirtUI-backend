package models

import (
	"encoding/json"
	"errors"
	"log"
	"virtui/api"
)

var clustersList []Cluster
var clusterGroupsList []ClusterGroup

type Cluster struct {
	Metadata struct {
		Architecture string `json:"architecture"`
		Config       struct {
			SchedulerInstance string `json:"scheduler.instance"`
		} `json:"config"`
		Database      bool     `json:"database"`
		Description   string   `json:"description"`
		FailureDomain string   `json:"failure_domain"`
		Groups        []string `json:"groups"`
		Message       string   `json:"message"`
		Roles         []string `json:"roles"`
		ServerName    string   `json:"server_name"`
		Status        string   `json:"status"`
		Url           string   `json:"url"`
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

//	func AddClusterAddress(cluster Cluster, group ClusterGroup) (string, error) {
//		GetClustersFromApi()
//		if clustersExist(cluster.Metadata.ServerName) {
//			return "", errors.New("Cluster already exists")
//		}
//		groups := group
//		groups.Members = append(groups.Members, cluster.Metadata.ServerName)
//		return api.Cli.Post("/1.0/cluster/groups", groups), nil
//	}
func CreateCluster(group ClusterGroup, clusters ...Cluster) (string, error) {
	for i := range clusters {
		if !clustersExist(clusters[i].Metadata.ServerName) {
			return "", errors.New("Cluster does not exists")
		}
		group.Members = append(group.Members, clusters[i].Metadata.ServerName)
	}
	return api.Cli.Post("/1.0/cluster/groups", group), nil
}

func DeleteCluster(serverName string) (string, error) {
	GetClustersFromApi()
	if clustersExist(serverName) {
		return api.Cli.Delete("/1.0/cluster/members/" + serverName), nil
	}
	return "", errors.New("Cluster does not exist")
}

func CreateContainerFromCluster(location, containerName string, fingerprint string) Operation {
	var operation Operation
	if !clustersExist(location) {
		log.Fatal("Cluster does not exist")
	}
	container := CreateContainer(containerName, fingerprint)
	container.Metadata.Location = location
	err := json.Unmarshal([]byte(api.Cli.Post("/1.0/instances", container)), &container)
	if err != nil {
		return Operation{}
	}
	return operation
}

func DeleteContainerFromCluster(location, containerName string) Operation {
	var operation Operation
	if !clustersExist(location) {
		log.Fatal("Cluster does not exist")
	}
	err := json.Unmarshal([]byte(api.Cli.Delete("/1.0/instances/"+containerName)), &operation)
	if err != nil {
		return Operation{}
	}
	return operation
}

func GetContainersFromCluster(clusterName string) ([]Container, error) {
	containersList := GetContainersFromApi()
	var containerListForCluster []Container
	for _, container := range containersList {
		if container.Metadata.Location == clusterName {
			containerListForCluster = append(containerListForCluster, container)
		}
	}
	return containerListForCluster, nil
}

func GetClustersFromApi() []Cluster {
	var clustersDetail []Cluster
	var clusterDetail Cluster
	var clusters clusters
	err := json.Unmarshal([]byte(api.Cli.Get("/1.0/cluster/members")), &clusters)
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

func GetClusterWithName(serverName string) (Cluster, error) {
	GetClustersFromApi()
	return clustersList[getIdClusterWithName(serverName)], nil
}

func GetClusterGroupWithName(groupName string) (ClusterGroup, error) {
	GetClustersFromApi()
	return clusterGroupsList[getIdClusterGroupWithName(groupName)], nil
}

func getIdClusterWithName(serverName string) int {
	for i, cluster := range clustersList {
		if cluster.Metadata.ServerName == serverName {
			return i
		}
	}
	return 0
}

func getIdClusterGroupWithName(groupName string) int {
	for i, group := range clusterGroupsList {
		if group.Name == groupName {
			return i
		}
	}
	return 0
}

func clustersExist(serverName string) bool {
	GetClustersFromApi()
	for _, cluster := range clustersList {
		if cluster.Metadata.ServerName == serverName {
			return true
		}
	}
	return false
}
