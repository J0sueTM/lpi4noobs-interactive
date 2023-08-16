package ui

import (
	"fyne.io/fyne/v2"
	"github.com/fyne-io/terminal"
)

func RunTTY(tty *terminal.Terminal) {
	_ = tty.RunLocalShell()
	fyne.CurrentApp().Quit()
}
