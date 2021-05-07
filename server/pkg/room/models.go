package room

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Participant struct {
	host bool
	conn *websocket.Conn
}

type RoomMap struct {
	mutex sync.RWMutex
	m     map[string][]Participant
}
