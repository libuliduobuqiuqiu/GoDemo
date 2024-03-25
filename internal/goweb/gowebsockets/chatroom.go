package gowebsockets

import (
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
	Name      string
	EventType int
	Message   string
}

var (
	chatUsers   = sync.Map{}
	subscribe   = make(chan Subscriber, 10)
	unsubscribe = make(chan string, 10)
	publish     = make(chan Event, 10)
)

func init() {
	startChatRoom()
}

func isUserExist(name string) error {
	if _, ok := chatUsers.Load(name); !ok {
		return fmt.Errorf("User: %s is exist.", name)
	}
	return nil
}

func broadcastMsg(event Event) {
	chatUsers.Range(func(key, value interface{}) bool {
		if key != event.Name {
			if user, ok := value.(Subscriber); ok {
				user.Conn.WriteMessage(event.EventType, []byte(event.Message))
			}
		}
		return true
	})
}

func startChatRoom() {
	for {
		select {
		case user := <-subscribe:
			if err := isUserExist(user.Name); err != nil {
				log.Println(err.Error())
			}

			chatUsers.Store(user.Name, user)
			publish <- Event{Name: user.Name, EventType: EventJoin, Message: fmt.Sprintf("%s User Join this chatroom.", user.Name)}

		case userName := <-unsubscribe:
			if err := isUserExist(userName); err == nil {
				log.Println(fmt.Errorf("User: %s is not exists.", userName).Error())
			}

			chatUsers.Delete(userName)
			publish <- Event{Name: userName, EventType: EventJoin, Message: fmt.Sprintf("%s User Leave this chatroom", userName)}

		case event := <-publish:
			broadcastMsg(event)
		}
	}
}
