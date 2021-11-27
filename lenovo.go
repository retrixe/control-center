package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/retrixe/control-panel/lenovo"
)

func GetLenovoSettings() *fyne.Container {
	conservationModeButton := widget.NewButton("Conservation Mode: N/A", func() {})
	if lenovo.IsConservationModeAvailable() {
		var updateConservationModeButton func()
		updateConservationModeButton = func() {
			if lenovo.IsConservationModeEnabled() {
				conservationModeButton.SetText("Conservation Mode: Enabled")
				conservationModeButton.OnTapped = func() {
					lenovo.SetConservationModeStatus(false)
					updateConservationModeButton()
				}
			} else {
				conservationModeButton.SetText("Conservation Mode: Disabled")
				conservationModeButton.OnTapped = func() {
					lenovo.SetConservationModeStatus(true)
					updateConservationModeButton()
				}
			}
		}
		go (func() {
			for {
				updateConservationModeButton()
				<-time.After(1 * time.Second)
			}
		})()
	} else {
		conservationModeButton.Disable()
	}

	titleLabel := widget.NewLabel("Lenovo")
	titleLabel.TextStyle.Bold = true
	return container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		conservationModeButton,
	)
}
