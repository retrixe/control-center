package main

import (
	"log"
	"os"

	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const version = "1.0.0-alpha.0"

// TODO: We need better elevation handling.

func main() {
	if len(os.Args) >= 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		println("control-panel version " + version)
		return
	}
	log.SetPrefix("[control-panel] ")

	a := app.New()
	w := a.NewWindow("Control Panel")

	settingsArea := GetLenovoSettings()
	var parent *fyne.Container
	parent = container.NewHBox(
		container.NewVBox(
			widget.NewButton("Lenovo", func() {
				parent.Remove(settingsArea)
				settingsArea = GetLenovoSettings()
				parent.Add(settingsArea)
			}),
			widget.NewButton("nouveau", func() {
				parent.Remove(settingsArea)
				settingsArea = getNouveauSettings()
				parent.Add(settingsArea)
			}),
		),
		widget.NewSeparator(),
		settingsArea,
	)
	w.SetContent(parent)

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}

func getNouveauSettings() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("nouveau: WIP"),
	)
}
