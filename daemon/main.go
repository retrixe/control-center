package main

import (
	"log"
	"os"

	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/godbus/dbus/v5"
)

const version = "1.0.0-alpha.0"

var conn *dbus.Conn

func InitialiseDBusConnection() {
	connection, err := dbus.ConnectSystemBus()
	if err != nil {
		log.Fatalln("Failed to connect to D-Bus system bus!", err)
	} else {
		log.Println("Successfully connected to D-Bus system bus!")
		conn = connection
	}
}

func main() {
	log.SetPrefix("[control-center] ")
	if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		log.Println("control-center version " + version)
		return
	} else if len(os.Args) == 2 && (os.Args[1] == "-d" || os.Args[1] == "--daemon") {
		InitialiseDBusConnection()
		defer conn.Close()
		StartDBusDaemon()
		return
	} else if len(os.Args) >= 2 {
		log.Println("Correct usage: ./control-center [-v or --version] [-d or --daemon]")
		return
	}

	InitialiseDBusConnection()
	defer conn.Close()

	a := app.New()
	w := a.NewWindow("Control Center")

	// TODO: Add a way to handle errors.
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
