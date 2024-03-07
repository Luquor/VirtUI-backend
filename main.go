package main

import (
	"virtui/models"
)

func main() {
	//test := models.CreateContainer("fklsdjlfdjsklfsdjk")
	//fmt.Println(test.Metadata.Description, test.Metadata.Id, test.Metadata.Status)
	//fmt.Println(models.GetOperationWithID(test.Metadata.Id))
	models.StartWebServer()
}
