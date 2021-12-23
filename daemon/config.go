package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/retrixe/control-center/daemon/lenovo"
)

type Config struct {
	LenovoConservationModeEnabled bool `json:"lenovoConservationModeEnabled"`
}

var config Config = Config{
	LenovoConservationModeEnabled: false,
}

func GetConfigPath() string {
	snapdir := os.Getenv("SNAP_DATA")
	path := "/etc/controlcenter/config.json"
	if snapdir != "" {
		path = filepath.Join(snapdir, "config.json")
	}
	return path
}

func WriteConfig() error {
	file, err := json.Marshal(&config)
	if err != nil {
		return err
	}

	err = os.WriteFile(GetConfigPath(), file, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

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
	if lenovo.IsConservationModeAvailable() {
		err := lenovo.SetConservationModeStatus(config.LenovoConservationModeEnabled)
		if err != nil {
			return err
		}
	}

	return nil
}
