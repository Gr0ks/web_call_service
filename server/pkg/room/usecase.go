package room

import (
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

func (r *RoomMap) Init() {
	r.m = make(map[string][]Participant)
}

func (r *RoomMap) Get(roomID string) []Participant {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.m[roomID]
}

func (r *RoomMap) CreateRoom() string {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i, _ := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.m[roomID] = []Participant{}

	return roomID
}

func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	p := Participant{host, conn}

	log.Println("Insert into Room with RoomID: ", roomID)
	r.m[roomID] = append(r.m[roomID], p)
}

func (r *RoomMap) DeleteRoom(roomID string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.m, roomID)
}
