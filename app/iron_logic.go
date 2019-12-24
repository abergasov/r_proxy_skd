package app

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	"time"
)

/**
Message from controller
*/
type IronMessage struct {
	Type     string
	Sn       int
	Messages []Message
}

type Message struct {
	Id            int     `json:"id"`
	Success       int     `json:"success,omitempty"`
	Operation     string  `json:"operation,omitempty"`
	ControllerIp  string  `json:"controller_ip,omitempty"`
	Active        int     `json:"active,omitempty"`
	Online        int     `json:"online,omitempty"`
	EventsSuccess *int    `json:"events_success,omitempty"`
	Granted       *int    `json:"granted,omitempty"`
	Card          string  `json:"card,omitempty"`
	Mode          *int    `json:"mode,omitempty"`
	Open          *int    `json:"open,omitempty"`
	OpenControl   *int    `json:"open_control,omitempty"`
	CloseControl  *int    `json:"close_control,omitempty"`
	Events        []Event `json:"events,omitempty"`
}

type Event struct {
	Flag  int
	Event int
	Time  string
	Card  string
}

/**
Response message to controller
*/
type Answer struct {
	Date     string    `json:"date"`
	Interval int       `json:"interval"`
	Messages []Message `json:"messages"`
}

/**
Parse message from controller and generate response
*/
func CatchMessage(jsonRow string) Answer {
	var result IronMessage
	err := json.Unmarshal([]byte(jsonRow), &result)
	if err != nil {
		LogMessage("Error parsing message", Warning, []string{jsonRow})
	}
	currentTime := time.Now()
	answer := Answer{
		Date:     currentTime.Format("2006-01-02 15:04:05"),
		Interval: 33,
		Messages: []Message{},
	}
	if len(result.Messages) > 0 {
		for i := 0; i < len(result.Messages); i++ {
			a, err := manageMessage(result.Messages[i], result.Sn)
			if err != nil {
				continue
			}
			answer.Messages = append(answer.Messages, a)
		}
	}
	return answer
}

/**
generate answer for single command
*/
func manageMessage(message Message, serialNumber int) (Message, error) {
	switch message.Operation {
	case "power_on":
		go NewDeviceRequest(strconv.Itoa(serialNumber), message.ControllerIp)
		return Message{
			Id:        rand.Intn(10000-1000) + 1000,
			Operation: "set_active",
			Active:    1,
			Online:    1,
		}, nil
	case "events":
		length := len(message.Events)
		return Message{
			Id:            message.Id,
			Operation:     "events",
			EventsSuccess: &length,
		}, nil
	case "check_access":
		granted := ManageDoor(strconv.Itoa(serialNumber), message.Card)
		return Message{
			Id:        message.Id,
			Operation: "check_access",
			Granted:   &granted,
		}, nil
	}
	if message.Success == 1 {
		go getConfirmCommand(serialNumber, message.Id)
	}
	return Message{}, errors.New("empty")
}

/**
confirm command execution from controller
*/
func getConfirmCommand(serial int, commandId int) {
	//todo
}
