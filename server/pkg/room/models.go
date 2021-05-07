package room

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Participant struct {
	host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	mutex sync.RWMutex
	M     map[string][]Participant
}
