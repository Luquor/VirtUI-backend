package main

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
	"virtui/models"

	"github.com/stretchr/testify/assert"
)

// TestCreationContainer
// for a valid return value.
func TestCreationContainer(t *testing.T) {
	name := "server"
	models.CreateContainer(name)
	cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+name)
	instances, err := cmd.Output()
	assert.Nil(t, err)
	assert.NotNil(t, instances)
}

func TestGetContainer(t *testing.T) {
	name := "server"
	models.CreateContainer(name)
	contai := models.GetContainerWithName(name).Metadata
	cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+contai.Name)
	instances, err := cmd.Output()
	assert.Nil(t, err)
	assert.NotNil(t, instances)
}
func TestSuppressionContainer(t *testing.T) {
	name := "server"
	models.CreateContainer(name)
	models.DeleteContainerWithName(name)
	supprimer := models.GetContainerWithName(name).Metadata
	fmt.Println("apres suppressions du conteneur:" + supprimer.Name)
	cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+name)
	instances, err := cmd.Output()
	if assert.NotNil(t, err) {
		var tab_byte []byte
		assert.Equal(t, string(tab_byte), string(instances))
	}

}

func list_lxd(nom string) string {
	cmd := exec.Command("curl", "-s", "-k", "--cert", "tls/client.crt", "--key", "tls/client.key", "-X", "GET", "https://127.0.0.1:8443/1.0/instances/"+nom)
	instances, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(instances)
}

// En attendant que le back soit fonctionnel j'ai fais cette fonction pour au moins tester mes test .
func creationContain(nom string) {
	cmd := exec.Command("curl", "-s", "-k", "--cert", "tls/client.crt", "--key", "tls/client.key", "-X", "POST", "https://127.0.0.1:8443/1.0/instances/"+nom)
	instances, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println(instances)
	}

}
func TestEtatContainer(t *testing.T) {
	//////verifier l'etat des differents conteneurs ///////
	//////start, stop, restart, freeze, unfreeze//////
}

///////////////////////////
