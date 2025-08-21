package main

import (
	"fmt"
	"golangR99/constants"
	"golangR99/controllers"
	"golangR99/view"
)

func main() {
	characters, psychubes := controllers.GetLists()
	view.Display(characters, psychubes)
	fmt.Println(constants.GET_CHARACTER_ENDPOINT("Medicine Pocket"))
}
