package main

import (
	"fyne.io/fyne/v2"
	"github.com/j0suetm/lpi4noobs-interactive/artifact"
	"github.com/j0suetm/lpi4noobs-interactive/ui"
)

func main() {
	ui_ := ui.New(800, 520)

	_, err := artifact.NewContent("content.json")
	if err != nil {
		fyne.LogError("content.go :: New", err)
		fyne.CurrentApp().Quit()
	}

	ui_.Window.ShowAndRun()
}
