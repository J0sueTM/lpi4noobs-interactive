package ui

import (
	"fyne.io/fyne/v2"
	fapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fyne-io/terminal"
)

type UI struct {
	App    fyne.App
	Window fyne.Window
}

func New(width float32, height float32) *UI {
	app := fapp.New()
	window := app.NewWindow("LPI4Noobs Interactive")
	window.Resize(fyne.NewSize(width, height))

	ttyConsole := terminal.New()

	sessionPage := container.NewHSplit(
		container.NewVSplit(
			widget.NewRichTextFromMarkdown("# FOOBAR\nentao familia vou falar sobre esse bgl aqui\n## Agua\nAgua eh muito bom ne"),
			widget.NewRichTextFromMarkdown("# Exercicio 1.1\nFaca um radiador em 2 passos..."),
		),
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(400, 540)),
			ttyConsole,
		),
	)

	window.SetContent(
		container.NewAppTabs(
			container.NewTabItem("Bem-Vindo", widget.NewButton("hello", func() {})),
			container.NewTabItem("Sessao 1", sessionPage),
		),
	)

	go func() {
		_ = ttyConsole.RunLocalShell()
		app.Quit()
	}()

	return &UI{
		App:    app,
		Window: window,
	}
}
