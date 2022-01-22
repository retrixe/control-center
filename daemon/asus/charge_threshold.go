package asus

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/retrixe/control-center/daemon/utils"
)

// https://wiki.archlinux.org/title/Laptop/ASUS

var ErrAsusChargeThresholdNotAvailable = errors.New(
	"asus charge threshold is not available on this system")

const ChargeThresholdSysFsFolder = "/sys/class/power_supply" // BAT0, BAT1, BATC and BATT
const ChargeThresholdSysFs = "/sys/class/power_supply/BAT0/charge_control_end_threshold"

func getChargeThresholdSysFs() string {
	return ChargeThresholdSysFs
}

func IsChargeThresholdAvailable() bool {
	modulesInfo, err := exec.Command("lsmod").Output()
	if err != nil {
		log.Println("Failed to run lsmod!")
		return false
	}
	modules := strings.Split(string(modulesInfo), "\n")
	for _, module := range modules {
		if strings.Fields(module)[0] == "asus-wmi" {
			// TODO: BAT0 can vary.
			_, err := os.ReadFile(getChargeThresholdSysFs())
			if os.IsNotExist(err) {
				return false
			} else if err != nil {
				log.Println("An unknown error occurred when checking for Asus charge threshold", err)
				return false
			}
			return true
		}
	}
	return false
}

func GetChargeThreshold() (int, error) {
	data, err := os.ReadFile(getChargeThresholdSysFs())
	if os.IsNotExist(err) {
		log.Println("Asus charge threshold status was checked despite no support for it", err)
		return 0, ErrAsusChargeThresholdNotAvailable
	} else if err != nil {
		log.Println("An unknown error occurred when checking for Asus charge threshold", err)
		return 0, err
	}
	chargeThreshold, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		log.Println("An unknown error occurred when checking for Asus charge threshold", err)
		return 0, err
	}
	return chargeThreshold, nil
}

func SetChargeThreshold(threshold int) error {
	if !IsChargeThresholdAvailable() { // Don't accidentally write to the file.
		return ErrAsusChargeThresholdNotAvailable
	}
	data := []byte(strconv.Itoa(threshold))
	err := utils.WriteFile(getChargeThresholdSysFs(), data)
	if err != nil {
		log.Println("Failed to set Asus charge threshold", err)
	}
	return err
}
