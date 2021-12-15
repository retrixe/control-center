package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/retrixe/control-center/daemon/nouveau"
)

// TODO: Complete D-Bus API and convert front-end to D-Bus API GUI for nouveau.

func GetNouveauSettings() *fyne.Container {
	devices, err := nouveau.NouveauGetDRIDevices()
	if err != nil {
		devices = make([]int, 0)
		log.Println("Failed to get devices using the nouveau driver", err)
	}

	var settings fyne.CanvasObject = widget.NewLabel("No devices with nouveau detected!")
	if len(devices) > 0 {
		settings = container.NewVBox()
		for _, device := range devices {
			// TODO: Add some meaningful settings.
			settings.(*fyne.Container).Add(widget.NewLabel("DRI device: " + strconv.Itoa(device)))
			settings.(*fyne.Container).Add(widget.NewSeparator())
		}
	}

	titleLabel := widget.NewLabel("nouveau")
	titleLabel.TextStyle.Bold = true
	return container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		settings,
	)
}
