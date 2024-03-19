package main

import (
	"virtui/models"
)

func main() {
	//test := models.CreateContainer("fklsdjlfdjsklfsdjk")
	//fmt.Println(test.Metadata.Description, test.Metadata.Id, test.Metadata.Status)
	//fmt.Println(models.GetOperationWithID(test.Metadata.Id))
	//test := models.GetConsoleForContainer("test")
	//fmt.Println(test.Status)
	//fmt.Println(test)
	//websocket := models.GetSocketsFromURLOperation(test.Operation)
	//fmt.Println(websocket)
	models.StartWebServer()
}
