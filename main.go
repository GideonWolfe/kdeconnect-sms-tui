package main

import (
	"github.com/godbus/dbus/v5"
	"log"
  "fmt"
)

func main() {

  // connect to the session dbus
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}

  // connect to the kde connect sms daemon
  obj := conn.Object("org.kde.kdeconnect", dbus.ObjectPath("/modules/kdeconnect/devices/51b84756d5ddeb87/sms"))

  // // register match rule to receive sms signals
  // See this for more funcs which create a match option https://godoc.org/github.com/godbus/dbus#MatchOption
	rule1 := dbus.WithMatchDestination("org.freedesktop.DBus.Properties.PropertiesChanged")
	fmt.Println(rule1)

	// Pass your match options to `AddMatchSignal`
	err = conn.AddMatchSignal(rule1)
	if err != nil {
		log.Fatal(err)
	}

  // asynchronous, no return value. Must subscribe to signals in conversation interface.
  foo := obj.Call("org.kde.kdeconnect.device.sms.requestAllConversations", 0); fmt.Printf("%#v", foo)
	if err != nil {
		log.Fatal(err)
	}
}
