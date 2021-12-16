package main

import (
	"log"
	"os"

	_ "embed"

	"github.com/godbus/dbus/v5"
)

const version = "1.0.0-alpha.0"

var conn *dbus.Conn

func InitialiseDBusConnection() {
	connection, err := dbus.ConnectSystemBus()
	if err != nil {
		log.Fatalln("Failed to connect to D-Bus system bus!", err)
	} else {
		log.Println("Successfully connected to D-Bus system bus!")
		conn = connection
	}
}

func main() {
	log.SetPrefix("[control-center] ")
	if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		log.Println("control-center version " + version)
		return
	} else if len(os.Args) >= 2 {
		log.Println("Correct usage: ./control-center [-v or --version]")
		return
	}

	InitialiseDBusConnection()
	defer conn.Close()
	StartDBusDaemon()
}
