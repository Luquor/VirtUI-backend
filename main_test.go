package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"virtui/models"

	"github.com/stretchr/testify/assert"
)

// TestCreationContainer
// for a valid return value.
func TestCreationContainer(t *testing.T) {
	cmd1 := exec.Command("sh", "-c", `lxc image list | grep -oP '^\| [^ALIAS|]*\s\| (\w*)' | sed 's/|.*| //'`)
	recupFingerPrint, err := cmd1.Output()
	name := "server"
	models.CreateContainer(name, strings.TrimSuffix(string(recupFingerPrint), "\n"))
	cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+name)
	instances, err := cmd.Output()
	//fmt.Println("cmd", cmd, "fin")
	//fmt.Println("instances,err", err, string(instances), "fin")
	assert.Nil(t, err)
	assert.NotNil(t, instances)
	exec.Command("lxc", "query", "--request", "DELETE", "/1.0/instances/"+name)
	//sudo lxc image copy images:f01555e462c4 didier:
	//afin de copier une image dans le cluster
}

func TestGetContainer(t *testing.T) {
	//name := "server2"
	//recupFingerPrint, err := exec.Command("lxc", "image", "list", "|", "grep", "-oP", `^\| [^ALIAS|]*\s\| (\w*)`, "|", " sed ", `s/|.*| //`).Output()
	//models.CreateContainer(name, string(recupFingerPrint))
	//contai := models.GetContainerWithName(name).Metadata
	//cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+contai.Name)
	//instances, err := cmd.Output()
	//fmt.Println(err, instances)
	//assert.Nil(t, err)
	//assert.NotNil(t, instances)
}
func TestSuppressionContainer(t *testing.T) {
	name := "server"
	//recupFingerPrint, err := exec.Command("lxc", "image", "list", "|", "grep", "-oP", `^\| [^ALIAS|]*\s\| (\w*)`, "|", " sed ", `s/|.*| //`).Output()
	//models.CreateContainer(name, string(recupFingerPrint))
	models.DeleteContainerWithName(name)
	//supprimer := models.GetContainerWithName(name).Metadata
	//fmt.Println("apres suppressions du conteneur:" + name)
	cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+name)
	instances, err := cmd.Output()
	fmt.Println(err, string(instances))
	//if assert.NotNil(t, err) {
	//var tab_byte []byte
	//assert.Equal(t, string(tab_byte), string(instances))
	//}

}

func TestEtatContainer(t *testing.T) {
	//////verifier l'etat des differents conteneurs ///////
	//////start, stop, restart, freeze, unfreeze//////
}

///////////////////////////S
