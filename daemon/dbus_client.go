// This file is unmaintained.
package main

import (
	"log"
	"strings"
)

const DBusCallErrorLogMessage = "Failed to call %s (is the control center daemon running?):"

const LenovoGetConservationModeStatusPath = "com.retrixe.ControlCenter.v0.LenovoGetConservationModeStatus"

func LenovoGetConservationModeStatus() (int, error) {
	var s int
	obj := conn.Object("com.retrixe.ControlCenter.v0", "/com/retrixe/ControlCenter/v0")
	err := obj.Call(LenovoGetConservationModeStatusPath, 0).Store(&s)
	if err != nil {
		log.Println(strings.Replace(DBusCallErrorLogMessage, "%s", LenovoGetConservationModeStatusPath, 1), err)
		return 0, err
	} else {
		return s, nil
	}
}

const LenovoSetConservationModePath = "com.retrixe.ControlCenter.v0.LenovoSetConservationMode"

func LenovoSetConservationMode(value bool) error {
	var s bool
	obj := conn.Object("com.retrixe.ControlCenter.v0", "/com/retrixe/ControlCenter/v0")
	err := obj.Call(LenovoSetConservationModePath, 0, value).Store(&s)
	if err != nil {
		log.Println(strings.Replace(DBusCallErrorLogMessage, "%s", LenovoGetConservationModeStatusPath, 1), err)
		return err
	} else {
		return nil
	}
}
