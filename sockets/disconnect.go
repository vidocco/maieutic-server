package sockets

import (
	"github.com/gorilla/websocket"
)

func Disconnect (conn *websocket.Conn) {
	room := clients[conn].Room
	if len(rooms[room].Users) <= 1 {
		rooms[room].Broadcast <- Event{ "disconnect", "" }
		<- rooms[room].Broadcast
		delete(rooms, room)
	}
	// TODO: disconnect user from room
	delete(clients, conn)
}