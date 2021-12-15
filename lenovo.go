package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetLenovoSettings() *fyne.Container {
	conservationModeButton := widget.NewButton("Conservation Mode: N/A", func() {})
	conservationModeStatus, err := LenovoGetConservationModeStatus()
	// TODO: Error handling.
	if err == nil && (conservationModeStatus == 0 || conservationModeStatus == 1) {
		var updateConservationModeButton func()
		updateConservationModeButton = func() {
			conservationModeStatus, _ := LenovoGetConservationModeStatus()
			if conservationModeStatus == 1 {
				conservationModeButton.SetText("Conservation Mode: Enabled")
				conservationModeButton.OnTapped = func() {
					LenovoSetConservationMode(false)
					updateConservationModeButton()
				}
			} else {
				conservationModeButton.SetText("Conservation Mode: Disabled")
				conservationModeButton.OnTapped = func() {
					LenovoSetConservationMode(true)
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
