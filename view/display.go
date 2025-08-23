package view

import (
	"fmt"
	"golangR99/constants"
	"golangR99/controllers"
	"golangR99/models"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// store all psychube data
var all_psychubes []models.Node

// track widget groups
var cGroup *widget.CheckGroup
var pGroup *widget.CheckGroup

// labels to keep updated
var selections *widget.Label
var results *widget.Label

func Display(characters, psychubes []models.Node) {
	a := app.New()
	w := a.NewWindow("Lists")
	// set all_psychubes
	all_psychubes = psychubes
	// set groups
	charScroll := groupToScroll(nodeToGroup(characters))
	psyScroll := groupToScroll(nodeToGroup(psychubes))
	// set containers
	top := container.NewGridWithColumns(2, charScroll, psyScroll)
	confirm, clear := setBottomElement()
	buttons := container.NewGridWithColumns(2, confirm, clear)
	// place results in a VScroll
	resultsScroll := container.NewVScroll(results)
	bottom := container.NewGridWithRows(3, selections, buttons, resultsScroll)
	// main container
	main := container.NewVSplit(top, bottom)
	main.Offset = 0.8
	// set to and run window
	w.SetContent(main)
	w.Resize(fyne.NewSize(800, 400))
	w.ShowAndRun()
}

func groupToScroll(group *widget.CheckGroup) *container.Scroll {
	return container.NewVScroll(group)
}

func nodeToGroup(nodes []models.Node) *widget.CheckGroup {
	if len(nodes) == 0 {
		return widget.NewCheckGroup(nil, nil)
	}

	nodeType := nodes[0].GetType()
	isCharacter := nodeType == constants.CHARACTER
	// set up options
	options := make([]string, len(nodes))
	for i, val := range nodes {
		options[i] = val.Name
	}

	group := widget.NewCheckGroup(options,
		func(selected []string) {
			setSelectedText()
		},
	)

	if isCharacter {
		cGroup = group
	} else {
		pGroup = group
	}

	return group
}

func setBottomElement() (*widget.Button, *widget.Button) {
	selections = widget.NewLabel("")
	confirm := widget.NewButton("Confirm", querySelections)
	clear := widget.NewButton("Clear", clearSelections)
	results = widget.NewLabel("")
	return confirm, clear
}

func querySelections() {
	result := controllers.GetAllConcurrently(cGroup.Selected, pGroup.Selected, all_psychubes)
	new_text := strings.Join(result, "\n")
	text := results.Text
	if len(text) == 0 {
		results.SetText(new_text)
	} else {
		results.SetText(fmt.Sprintf("%s\n%s", text, new_text))
	}
	clearSelections()
}

func setSelectedText() {
	var characters, psychubes string
	if len(cGroup.Selected) == 0 {
		characters = ""
	} else {
		characters = "Characters: " + strings.Join(cGroup.Selected, ", ")
	}
	if len(pGroup.Selected) == 0 {
		psychubes = ""
	} else {
		psychubes = "Psychubes: " + strings.Join(pGroup.Selected, ", ")
	}
	selections.SetText(fmt.Sprintf("%s\n%s", characters, psychubes))
}

func clearSelections() {
	cGroup.SetSelected([]string{})
	pGroup.SetSelected([]string{})
	setSelectedText()
}
