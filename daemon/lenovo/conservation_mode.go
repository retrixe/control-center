package lenovo

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

var ErrLenovoConservationModeNotAvailable = errors.New(
	"lenovo conservation mode is not available on this system")

const ConservationModeSysFs = "/sys/bus/platform/drivers/ideapad_acpi/VPC2004:00/conservation_mode"

func IsConservationModeAvailable() bool {
	modulesInfo, err := exec.Command("lsmod").Output()
	if err != nil {
		log.Println("Failed to run lsmod!")
		return false
	}
	modules := strings.Split(string(modulesInfo), "\n")
	for _, module := range modules {
		if strings.Fields(module)[0] == "ideapad_laptop" {
			// TODO: VPC2004:00 can vary.
			_, err := os.ReadFile(ConservationModeSysFs)
			if os.IsNotExist(err) {
				return false
			} else if err != nil {
				log.Println("An unknown error occurred when checking for Lenovo conservation mode", err)
				return false
			}
			return true
		}
	}
	return false
}

func IsConservationModeEnabled() bool {
	data, err := os.ReadFile(ConservationModeSysFs)
	if os.IsNotExist(err) {
		log.Println("Lenovo conservation mode status was checked despite no support for it", err)
		return false
	} else if err != nil {
		log.Println("An unknown error occurred when checking for Lenovo conservation mode", err)
		return false
	}
	return string(data) == "1\n"
}

func SetConservationModeStatus(mode bool) error {
	if !IsConservationModeAvailable() { // Don't accidentally write to the file.
		return ErrLenovoConservationModeNotAvailable
	}
	data := []byte("0")
	if mode {
		data = []byte("1")
	}
	err := os.WriteFile(ConservationModeSysFs, data, os.ModePerm)
	if err != nil {
		log.Println("Failed to set Lenovo conservation mode", err)
	}
	return err
}
