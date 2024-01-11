package main

import (
	"fmt"
	"virtui/models"
)

func main() {
	test := models.CreateContainer("JeSuisLUCAS")
	fmt.Println(test.Metadata.Description, test.Metadata.Id, test.Metadata.Status)
	lastoperation := models.GetLastOperation()
	fmt.Println(lastoperation.Metadata.Description, lastoperation.Metadata.Id, lastoperation.Metadata.Status)
	models.StartWebServer()
}
