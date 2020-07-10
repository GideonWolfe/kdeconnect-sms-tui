package main

import (
	"github.com/godbus/dbus"
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
  obj := conn.Object("org.kde.kdeconnect", //Well known name on the bus (`busctl list` will show these)
  // dbus.ObjectPath("/modules/kdeconnect/devices/51b84756d5ddeb87/sms")) //Object path (`busctl tree <well known name>` shows these)
  dbus.ObjectPath("/modules/kdeconnect/devices/51b84756d5ddeb87/sms")) //Object path (`busctl tree <well known name>` shows these)

  foo := obj.Call("org.kde.kdeconnect.device.sms.requestAllConversations", 0); fmt.Printf("%#v", foo)
	if err != nil {
		log.Fatal(err)
	}
}
