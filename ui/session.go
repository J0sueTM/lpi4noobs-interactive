package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fyne-io/terminal"
)

type SessionUI struct {
	Content   *widget.RichText
	Exercise  *widget.RichText
	TTY       *terminal.Terminal
	Container *container.Split
}

func NewSession() *SessionUI {
	content := widget.NewRichTextFromMarkdown("")
	exercise := widget.NewRichTextFromMarkdown("")
	tty := terminal.New()
	container := container.NewHSplit(
		container.NewVSplit(content, exercise),
		tty,
	)

	return &SessionUI{
		Content:   content,
		Exercise:  exercise,
		TTY:       tty,
		Container: container,
	}
}
