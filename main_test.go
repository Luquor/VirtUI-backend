package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"
	"virtui/models"

	"github.com/stretchr/testify/assert"
)

// TestCreationContainer
// for a valid return value.
func TestCreationContainer(t *testing.T) {
	cmd1 := exec.Command("sh", "-c", `lxc image list | grep -oP '^\| [^ALIAS|]*\s\| (\w*)' | sed 's/|.*| //'`)
	recupFingerPrint, _ := cmd1.Output()
	name := "server"
	models.CreateContainer(name, strings.TrimSuffix(string(recupFingerPrint), "\n"), "")
	cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+name)
	instances, err := cmd.Output()
	//
	//	assert.Nil(t, err)
	//	assert.NotNil(t, instances)

	//	fmt.Println("cmd",err, "fin")
	//fmt.Println("instances,err", err, string(instances), "fin")
	assert.Nil(t, err)
	assert.NotNil(t, instances)
	exec.Command("lxc", "query", "--request", "DELETE", "/1.0/instances/"+name)

}

//	func TestGetContainer(t *testing.T) {
//		//name := "server2"
//		//recupFingerPrint, err := exec.Command("lxc", "image", "list", "|", "grep", "-oP", `^\| [^ALIAS|]*\s\| (\w*)`, "|", " sed ", `s/|.*| //`).Output()
//		//models.CreateContainer(name, string(recupFingerPrint))
//		//contai := models.GetContainerWithName(name).Metadata
//		//cmd := exec.Command("lxc", "query", "--request", "GET", "/1.0/instances/"+contai.Name)
//		//instances, err := cmd.Output()
//		//fmt.Println(err, instances)
//		//assert.Nil(t, err)
//		//assert.NotNil(t, instances)
//	}
func TestSuppressionContainer(t *testing.T) {
	name := "server"

	cmd1 := exec.Command("sh", "-c", `lxc image list | grep -oP '^\| [^ALIAS|]*\s\| (\w*)' | sed 's/|.*| //'`)
	recupFingerPrint, _ := cmd1.Output()
	models.CreateContainer(name, strings.TrimSuffix(string(recupFingerPrint), "\n"), "")
	time.Sleep(30 * time.Second)
	models.DeleteContainerWithName(name)
	//supprimer := models.GetContainerWithName(name).Metadata
	//	exec.Command("lxc", "query", "--request", "DELETE", "/1.0/instances/"+name)
	time.Sleep(10 * time.Second)
	cmd := exec.Command("sh", "-c", `lxc list | grep server`)
	instances, err := cmd.Output()
	fmt.Println(err)
	assert.Equal(t, string(instances), "")

}

func TestEtatContainer(t *testing.T) {
	//////verifier l'etat des differents conteneurs ///////
	//////start, stop, restart, freeze, unfreeze//////
}

///////////////////////////S
