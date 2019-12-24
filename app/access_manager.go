package app

import (
	"fmt"
)

type Device struct {
	Serial     string
	CwId       string
	DeviceName string
}

var deviceMap []Device

/**
notify about new device
*/
func NewDeviceRequest(serial string, ipAddress string) {
	device := findDevice(serial)
	LogMessage("New device request catches", Info, []string{device.DeviceName, "IP " + ipAddress})
}

/**
check user can open door
*/
func ManageDoor(device string, card string) int {
	//todo
	grantAccess := 0
	go logAccess(card, device, grantAccess == 1)
	return grantAccess
}

/**
write log open/close door for card
*/
func logAccess(card string, room string, grantAccess bool) bool {
	//todo
	if !grantAccess {
		notifyAccess(card, room)
	}
	return true
}

func findDevice(serial string) Device {
	for i := range deviceMap {
		if deviceMap[i].Serial == serial {
			return deviceMap[i]
		}
	}
	return Device{DeviceName: "Unknown device serial " + serial, CwId: "0", Serial: serial}
}

func notifyAccess(card string, room string) {
	msg := fmt.Sprintf("Попытка открыть %s\nкарта %s", card, room)
	chatId := conf.Chat
	TelegramMessage(chatId, msg, "html")
}
