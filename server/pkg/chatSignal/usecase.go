package chatSignal

import (
	"encoding/json"
	"log"
	"net/http"
)

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.CreateRoom()

	log.Println(AllRooms.M)
	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

func broadcaster() {
	for {
		msg := <-broadcast

		for _, client := range AllRooms.M[msg.RoomID] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)

				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
				}
			}
		}
	}
}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomID, ok := r.URL.Query()["roomID"]

	if !ok {
		log.Println("roomID missing in URL Parameters")
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}

	AllRooms.InsertIntoRoom(roomID[0], false, ws)

	go broadcaster()

	for {
		var msg broadcastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Println("Reed Error: ", err)
			return
		}

		msg.Client = ws
		msg.RoomID = roomID[0]

		log.Println(msg.Message)

		broadcast <- msg
	}
}
