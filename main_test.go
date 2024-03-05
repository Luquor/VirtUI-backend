package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

// TestCreationContainer
// for a valid return value.
func TestCreationContainer2(t *testing.T) {
	name := "server"
	//models.CreateContainer(name)
	msg := list_lxd(name)
	fmt.Println(string(msg))
	tab_str := strings.Split(msg, ":")
	//fmt.Println(tab_str)

}

func list_lxd(nom string) string {
	cmd := exec.Command("curl", "-s", "-k", "--cert", "tls/client.crt", "--key", "tls/client.key", "-X", "GET", "https://127.0.0.1:8443/1.0/instances/"+nom)
	instances, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(instances)
}
//En attendant que le back soit fonctionnel j'ai fais cette fonction pour au moins tester mes test .
func creationContain(nom string) {
	cmd := exec.Command("curl", "-s", "-k", "--cert", "tls/client.crt", "--key", "tls/client.key", "-X", "POST", "https://127.0.0.1:8443/1.0/instances/"+nom)
	instances, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}

///////////////////////////
