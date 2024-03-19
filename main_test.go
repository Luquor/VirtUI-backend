package main

import (
	"fmt"
	"os/exec"
	"testing"
	"virtui/models"
)

// TestCreationContainer
// for a valid return value.
func TestCreationContainer(t *testing.T) {
	cmd1 := exec.Command("lxc", "image", "list", "|", "grep", "-oP", `^\| [^ALIAS|]*\s\| (\w*)`, "|", " sed ", `s/|.*| //`)
	recupFingerPrint, err := cmd1.Output()
	fmt.Println(recupFingerPrint)
	name := "server"
	models.CreateContainer(name, string(recupFingerPrint))
	cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+name)
	instances, err := cmd.Output()
	fmt.Println(err, instances)
	//assert.Nil(t, err)
	//assert.NotNil(t, instances)
	//sudo lxc image copy images:f01555e462c4 didier:
	//afin de copier une image dans le cluster
}

func TestGetContainer(t *testing.T) {
	name := "server"
	recupFingerPrint, err := exec.Command("lxc", "image", "list", "|", "grep", "-oP", `^\| [^ALIAS|]*\s\| (\w*)`, "|", " sed ", `s/|.*| //`).Output()
	models.CreateContainer(name, string(recupFingerPrint))
	contai := models.GetContainerWithName(name).Metadata
	cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+contai.Name)
	instances, err := cmd.Output()
	fmt.Println(err, instances)
	//assert.Nil(t, err)
	//assert.NotNil(t, instances)
}
func TestSuppressionContainer(t *testing.T) {
	name := "server"
	recupFingerPrint, err := exec.Command("lxc", "image", "list", "|", "grep", "-oP", `^\| [^ALIAS|]*\s\| (\w*)`, "|", " sed ", `s/|.*| //`).Output()
	models.CreateContainer(name, string(recupFingerPrint))
	models.DeleteContainerWithName(name)
	supprimer := models.GetContainerWithName(name).Metadata
	fmt.Println("apres suppressions du conteneur:" + supprimer.Name)
	cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+name)
	instances, err := cmd.Output()
	fmt.Println(err, instances)
	//if assert.NotNil(t, err) {
		//var tab_byte []byte
		//assert.Equal(t, string(tab_byte), string(instances))
	//}

}

func TestEtatContainer(t *testing.T) {
	//////verifier l'etat des differents conteneurs ///////
	//////start, stop, restart, freeze, unfreeze//////
}

///////////////////////////
