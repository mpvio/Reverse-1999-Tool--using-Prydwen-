package view

import (
	"fmt"
	"golangR99/constants"
	"golangR99/models"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var selected_characters []string
var selected_psychubes []string

func Display(characters, psychubes []models.Node) {
	a := app.New()
	w := a.NewWindow("Lists")
	// set groups
	charScroll := groupToScroll(nodeToGroup(characters))
	psyScroll := groupToScroll(nodeToGroup(psychubes))
	// set to container
	top := container.NewGridWithColumns(2, charScroll, psyScroll)
	w.SetContent(top)
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
		},
	)

	return group
}
