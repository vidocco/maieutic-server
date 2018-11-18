package sockets

import (
	"github.com/gorilla/websocket"
)

func Disconnect (conn *websocket.Conn) {
	room := clients[conn].Room
	if len(rooms[room].Users) <= 1 {
		rooms[room].Broadcast <- message{ "disconnect", content { "" } }
		<- rooms[room].Broadcast
		delete(rooms, room)
		delete(keys, room)
	}
	delete(clients, conn)
}