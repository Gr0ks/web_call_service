package main

import (
	"log"
	"net/http"
	"web_call_service/server/pkg/chatSignal"
)

func main() {
	chatSignal.AllRooms.Init()

	http.HandleFunc("/create", chatSignal.CreateRoomRequestHandler)
	http.HandleFunc("/join", chatSignal.JoinRoomRequestHandler)

	log.Println("Starting Server on Port 8088")
	err := http.ListenAndServeTLS(":8088", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
