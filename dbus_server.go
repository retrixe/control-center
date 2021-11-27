package main

import (
	"log"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/retrixe/control-panel/lenovo"
)

const intro = introspect.IntrospectDeclarationString + `
<node>
	<interface name="com.retrixe.ControlPanel.v0">
		<method name="LenovoGetConservationModeStatus">
			<arg direction="out" type="i"/>
		</method>
		<method name="LenovoSetConservationMode">
		  <arg direction="in" type="u"/>
			<arg direction="out" type="u"/>
		</method>
	</interface>
` + introspect.IntrospectDataString + `</node>`

type api string

func (a api) LenovoGetConservationModeStatus() (int, *dbus.DBusError) {
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

func (a api) LenovoSetConservationMode(status uint) (uint, *dbus.DBusError) {
	if lenovo.SetConservationModeStatus(status == 1) {
		return 1, nil
	}
	return 0, nil
}

func StartDBusDaemon() {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		log.Fatalln("Failed to connect to D-Bus system bus!", err)
	}

	f := api("ControlPanel v0 API")
	conn.Export(f, "/com/retrixe/ControlPanel/v0", "com.retrixe.ControlPanel.v0")
	conn.Export(introspect.Introspectable(intro), "/com/retrixe/ControlPanel/v0",
		"org.freedesktop.DBus.Introspectable")

	reply, err := conn.RequestName("com.retrixe.ControlPanel.v0", dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Fatalln("Failed to request D-Bus name com.retrixe.ControlPanel.v0", err)
	} else if reply != dbus.RequestNameReplyPrimaryOwner {
		log.Fatalln("D-Bus name com.retrixe.ControlPanel.v0 already taken")
	}

	log.Println("Listening on D-Bus name com.retrixe.ControlPanel.v0.")
	select {}
}
