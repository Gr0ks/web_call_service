package signal

import (
	"net/http"
	"web_call_service/server/pkg/room"

	"github.com/gorilla/websocket"
)

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

var (
	AllRooms room.RoomMap
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	broadcast = make(chan broadcastMsg)
)
