package main

import (
	"fmt"
	"golangR99/controllers"
)

func main() {
	_, psychubes := controllers.GetLists()
	// view.Display(characters, psychubes)
	controllers.GetCharacter("alexios")
	controllers.GetCharacter("apple")
	fmt.Println(controllers.GetPsychube("As the Bell Tolls", psychubes))
}
