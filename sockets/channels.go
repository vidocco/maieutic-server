package sockets

import "github.com/gorilla/websocket"

type Event struct {
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

type QueuedEvent struct {
	Event Event
	Client *websocket.Conn
}

type room struct {
	Users []*websocket.Conn
	Broadcast chan Event
}

type user struct {
	Client *websocket.Conn
	Room string
}

var Broadcast = make(chan QueuedEvent)
var rooms = make(map[string] *room)
var clients = make(map[*websocket.Conn] *user)

func newRoom (user *websocket.Conn) *room {
	return &room{
		[]*websocket.Conn{user},
		make(chan Event),
	}
}

func newUser (client *websocket.Conn, roomId string) *user {
	return &user{
		client,
		roomId,
	}
}

func addRoom (id string, user *websocket.Conn) {
	/**
	 * TODO: when creating a new room, the user needs a room name and a server password (which will be used in a middleware).

	- we check if the password matches the server pass.
	- we use the room name to create a new room in the rooms map.
	- return success to the client.

	2) Send one message encrypted with a single key, said key is encrypted multiple times (with the public keys of each user) and sent.
	 */
	rooms[id] = newRoom(user)
	clients[user].Room = id
	go roomHandler(id)
}

func joinRoom (id string, user *websocket.Conn) {
	rooms[id].Users = append(rooms[id].Users, user)
	clients[user].Room = id
}

func addUser (client *websocket.Conn) {
	clients[client] = newUser(client, "")
}