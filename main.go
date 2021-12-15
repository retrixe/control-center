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

// TODO: Complete D-Bus API and convert front-end to D-Bus API GUI.

func main() {
	log.SetPrefix("[control-center] ")
	if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		println("control-center version " + version)
		return
	} else if len(os.Args) == 2 && (os.Args[1] == "-d" || os.Args[1] == "--daemon") {
		StartDBusDaemon()
		return
	} else if len(os.Args) >= 2 {
		println("Correct usage: ./control-center [-v or --version] [-d or --daemon]")
		return
	}

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
				settingsArea = GetNouveauSettings()
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
