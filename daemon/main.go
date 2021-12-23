package main

import (
	"log"
	"os"

	_ "embed"

	"github.com/godbus/dbus/v5"
)

const version = "1.0.0-alpha.0"

var conn *dbus.Conn
var errLog = log.New(os.Stderr, "[control-center] ", log.LstdFlags)

func InitialiseDBusConnection() {
	connection, err := dbus.ConnectSystemBus()
	if err != nil {
		errLog.Fatalln("Failed to connect to D-Bus system bus!", err)
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

	err := LoadConfig()
	if os.IsNotExist(err) {
		log.Println("Creating new config file at " + GetConfigPath() + ".")
		err = WriteConfig()
		if err != nil {
			log.Println("Failed to write persistent config file! "+
				"Any changes made during this session may be lost on reboot!", err)
		}
	} else if err != nil {
		errLog.Fatalln("Failed to read config file for unknown reasons!", err)
	}

	err = ApplyConfig()
	if err != nil {
		log.Println("Failed to apply config! Some settings may not have been applied.", err)
	}

	StartDBusDaemon()
}
