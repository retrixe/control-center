package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	//	LenovoConservationModeEnabled bool `json:"lenovoConservationModeEnabled"`
}

var config Config = Config{
	//	LenovoConservationModeEnabled: false,
}

func GetConfigPath() string {
	snapdir := os.Getenv("SNAP_DATA")
	path := "/etc/controlcenter/config.json"
	if snapdir != "" {
		path = filepath.Join(snapdir, "config.json")
	}
	return path
}

func SaveConfig() {
	file, err := json.Marshal(&config)
	if err != nil {
		log.Println("Failed to write persistent config file! "+
			"Any changes made during this session may be lost on reboot!", err)
	}

	err = os.WriteFile(GetConfigPath(), file, os.ModePerm)
	if err != nil {
		log.Println("Failed to write persistent config file! "+
			"Any changes made during this session may be lost on reboot!", err)
	}
}

// TODO: ReadSystemConfig if config does not exist? Cross-OS compat needs to be better.

func LoadConfig() error {
	file, err := os.ReadFile(GetConfigPath())
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}
	return nil
}

func ApplyConfig() error {
	// Lenovo Conservation Mode.
	// if lenovo.IsConservationModeAvailable() {
	// 	err := lenovo.SetConservationModeStatus(config.LenovoConservationModeEnabled)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
