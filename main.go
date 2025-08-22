package main

import (
	"golangR99/controllers"
)

func main() {
	_, psychubes := controllers.GetLists()
	// view.Display(characters, psychubes)
	alexios := controllers.GetCharacter("alexios")
	apple := controllers.GetCharacter("apple")
	bell := controllers.GetPsychube("As the Bell Tolls", psychubes)
	// file writing test
	controllers.WriteToFile(alexios.Convert())
	controllers.WriteToFile(apple.Convert())
	controllers.WriteToFile(bell)
}
