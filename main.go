package main

import (
	"golangR99/controllers"
	"golangR99/view"
)

func main() {
	characters, psychubes := controllers.GetLists()
	view.Display(characters, psychubes)
	controllers.GetCharacter("sonetto")
	controllers.GetCharacter("ezio auditore")
}
