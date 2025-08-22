package main

import (
	"golangR99/controllers"
)

func main() {
	_, psychubes := controllers.GetLists()
	// view.Display(characters, psychubes)
	alexios := controllers.GetCharacter("alexios")
	apple := controllers.GetCharacter("apple")
	ezio := controllers.GetCharacter("ezio auditore")
	bell := controllers.GetPsychube("As the Bell Tolls", psychubes)
	// file writing test
	controllers.WriteToFile(alexios)
	controllers.WriteToFile(apple)
	controllers.WriteToFile(ezio)
	controllers.WriteToFile(bell)
}
