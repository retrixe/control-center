package main

import (
	"log"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/retrixe/control-center/daemon/lenovo"
	"github.com/retrixe/control-center/daemon/nouveau"
)

// https://dbus.freedesktop.org/doc/dbus-specification.html
// https://dbus.freedesktop.org/doc/dbus-tutorial.html

// TODO: Restrict API to wheel/sudo users.

const intro = introspect.IntrospectDeclarationString + `
<node>
	<interface name="com.retrixe.ControlCenter.v0">
		<method name="LenovoGetConservationModeStatus">
			<arg direction="out" type="n"/>
		</method>
		<method name="LenovoSetConservationMode">
		  <arg direction="in" type="b"/>
		</method>

		<method name="NouveauGetDRIDevices">
			<arg direction="out" type="ai"/>
		</method>
		<method name="NouveauSetPowerState">
			<arg direction="in" type="i"/>
			<arg direction="in" type="s"/>
		</method>
	</interface>` + introspect.IntrospectDataString + `</node>`

func StartDBusDaemon() {
	f := DBusAPI("Control Center v0 API")
	conn.Export(f, "/com/retrixe/ControlCenter/v0", "com.retrixe.ControlCenter.v0")
	conn.Export(introspect.Introspectable(intro), "/com/retrixe/ControlCenter/v0",
		"org.freedesktop.DBus.Introspectable")

	reply, err := conn.RequestName("com.retrixe.ControlCenter.v0", dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatalln("Failed to request D-Bus name com.retrixe.ControlCenter.v0", err)
	} else if reply != dbus.RequestNameReplyPrimaryOwner {
		log.Fatalln("D-Bus name com.retrixe.ControlCenter.v0 already taken")
	}

	stdLog.Println("Listening on D-Bus name com.retrixe.ControlCenter.v0.")
	select {}
}

type DBusAPI string

func (f DBusAPI) LenovoGetConservationModeStatus() (int16, *dbus.Error) {
	if lenovo.IsConservationModeAvailable() {
		if lenovo.IsConservationModeEnabled() {
			return 1, nil
		} else {
			return 0, nil
		}
	} else {
		return -1, nil
	}
}

func (f DBusAPI) LenovoSetConservationMode(status bool) *dbus.Error {
	if err := lenovo.SetConservationModeStatus(status); err != nil {
		return &dbus.Error{
			Name: "A daemon error occurred",
			Body: []interface{}{err.Error()},
		}
	}
	config.LenovoConservationModeEnabled = status
	SaveConfig()
	return nil
}

func (f DBusAPI) NouveauGetDRIDevices() ([]int, *dbus.Error) {
	devices, err := nouveau.NouveauGetDRIDevices()
	if err != nil {
		return nil, &dbus.Error{
			Name: "A daemon error occurred",
			Body: []interface{}{err.Error()},
		}
	}
	return devices, nil
}

// TODO: NouveauGetPowerStates

func (f DBusAPI) NouveauSetPowerState(driDevice int, state string) *dbus.Error {
	err := nouveau.NouveauSetPowerState(driDevice, state)
	if err != nil {
		return &dbus.Error{
			Name: "A daemon error occurred",
			Body: []interface{}{err.Error()},
		}
	}
	return nil
}
