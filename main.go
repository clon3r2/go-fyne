package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type App struct {
	label *widget.Label
}

func (app *App) btnAction() (*widget.Label, *widget.Entry, *widget.Entry, *widget.Button) {
	label := widget.NewLabel("hello world")
	entry1 := widget.NewEntry()
	entry2 := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		num1, err := strconv.Atoi(entry1.Text)
		if err != nil {
			return
		}
		num2, err := strconv.Atoi(entry2.Text)
		if err != nil {
			return
		}

		app.label.SetText(strconv.Itoa(num1 * num2))
	})
	app.label = label

	return app.label, entry1, entry2, btn
}

func main() {
	externalApp := app.New()
	w := externalApp.NewWindow("MyWindow")
	internalApp := App{}

	label, entry1, entry2, btn := internalApp.btnAction()
	conH := container.NewHBox(entry1, entry2)
	conH.Resize(fyne.Size{
		Width: 350,
	})
	w.SetContent(container.NewVBox(label, conH, btn))
	w.Resize(fyne.Size{
		Width:  350,
		Height: 500,
	})
	w.ShowAndRun()
}
