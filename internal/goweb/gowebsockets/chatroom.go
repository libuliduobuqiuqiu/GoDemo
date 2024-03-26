package gowebsockets

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	EventJoin = iota
	EventLeave
	EventMessage
)

type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

type Event struct {
	Name      string `json:"name"`
	EventType int    `json:"event_type"`
	Message   string `json:"message"`
}

var (
	chatUsers   = sync.Map{}
	subscribe   = make(chan Subscriber, 10)
	unsubscribe = make(chan string, 10)
	publish     = make(chan Event, 10)
)

func init() {
	go startChatRoom()
}

func join(s Subscriber) {
	subscribe <- s
}

func leave(name string) {
	unsubscribe <- name
}

func sendMessage(event Event) {
	publish <- event
}

func isUserExist(name string) error {
	if _, ok := chatUsers.Load(name); ok {
		return fmt.Errorf("User: %s is exist.", name)
	}
	return nil
}

func broadcastMsg(event Event) {
	fmt.Printf("%s: %s\n", event.Name, event.Message)

	switch event.EventType {
	case EventJoin, EventLeave:
		fmt.Println(event.Message)
	case EventMessage:
		fmt.Printf("%s: %s\n", event.Name, event.Message)
	}

	chatUsers.Range(func(key, value interface{}) bool {
		if key != event.Name {
			if user, ok := value.(Subscriber); ok {
				msg, err := json.Marshal(event)
				if err != nil {
					return false
				}
				user.Conn.WriteMessage(websocket.TextMessage, msg)
			}
		}
		return true
	})
}

func startChatRoom() {
	for {
		select {
		case user := <-subscribe:
			chatUsers.Store(user.Name, user)
			publish <- Event{Name: user.Name, EventType: EventJoin, Message: fmt.Sprintf("User: %s Join this chatroom.", user.Name)}

		case userName := <-unsubscribe:
			if err := isUserExist(userName); err == nil {
				log.Println(fmt.Errorf("User: %s is not exists.", userName).Error())
			}

			chatUsers.Delete(userName)
			publish <- Event{Name: userName, EventType: EventJoin, Message: fmt.Sprintf("User: %s Leave this chatroom", userName)}

		case event := <-publish:
			broadcastMsg(event)
		}
	}
}
