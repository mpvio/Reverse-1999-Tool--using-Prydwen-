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

// track selected items
var selected_characters []string
var selected_psychubes []string

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
	button := setBottomElement()
	bottom := container.NewGridWithRows(3, selections, button, results)
	// main container
	main := container.NewVSplit(top, bottom)
	main.Offset = 0.6
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
	_ = nodeType == constants.CHARACTER
	// set up options
	options := make([]string, len(nodes))
	for i, val := range nodes {
		options[i] = val.Name
	}

	group := widget.NewCheckGroup(options,
		func(selected []string) {
			if nodeType == constants.CHARACTER {
				selected_characters = selected
				fmt.Println(selected_characters)
			} else {
				selected_psychubes = selected
				fmt.Println(selected_psychubes)
			}
			setSelectedText()
		},
	)

	return group
}

func setBottomElement() *widget.Button {
	selections = widget.NewLabel("")
	confirm := widget.NewButton("Confirm", querySelections)
	results = widget.NewLabel("")
	return confirm
}

func querySelections() {
	result := controllers.GetAll(selected_characters, selected_psychubes, all_psychubes)
	results.SetText(strings.Join(result, "\n"))
}

func setSelectedText() {
	var characters, psychubes string
	if len(selected_characters) == 0 {
		characters = ""
	} else {
		characters = "Characters:" + strings.Join(selected_characters, ", ")
	}
	if len(selected_psychubes) == 0 {
		psychubes = ""
	} else {
		psychubes = "Psychubes:" + strings.Join(selected_psychubes, ", ")
	}
	selections.SetText(fmt.Sprintf("%s\n%s", characters, psychubes))
}
