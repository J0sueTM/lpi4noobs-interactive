package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	fapp "fyne.io/fyne/v2/app"
)

type UI struct {
	App     fyne.App
	Window  fyne.Window
	Session *SessionUI
}

const (
	appName    string = "LPI4Noobs Interactive"
	appVersion string = "0.0.1"
)

func New(width float32, height float32) *UI {
	app := fapp.New()

	windowTitle := fmt.Sprintf("%s %s", appName, appVersion)
	window := app.NewWindow(windowTitle)
	window.Resize(fyne.NewSize(width, height))

	session := NewSession()

	window.SetContent(session.Container)

	go RunTTY(session.TTY)

	return &UI{
		App:     app,
		Window:  window,
		Session: session,
	}
}
