package sockets

import "github.com/gorilla/websocket"

type Event struct {
	Type string
	Content interface{}
}

type content struct {
	Encrypted string `json:"encrypted"`
}

type message struct {
	Type string
	Content content
}

type room struct {
	Users []*websocket.Conn
	Broadcast chan message
}

type roomKey struct {
	Public string
	Private string
}

type user struct {
	Client *websocket.Conn
	Room string
}

var Broadcast = make(chan Event)
var rooms = make(map[string]room)
var keys = make(map[string]roomKey)
var clients = make(map[*websocket.Conn]user)

func newRoom (user *websocket.Conn) room {
	return room{
		[]*websocket.Conn{user},
		make(chan message),
	}
}

func newKeys (key string) roomKey {
	public, private := genKeys()
	return roomKey{
		public,
		private,
	}
}

func genKeys() (public string, private string) {
	// TODO: generate a public and private key and encrypt with specified key
	return "", ""
}

func newUser (client *websocket.Conn, roomId string) user {
	return user{
		client,
		roomId,
	}
}

func addRoom (id string, key string, user *websocket.Conn) {
	/**
	 * TODO: when creating a new room, the user needs a room name, password and key.
	 * the room name is used to add to the rooms map. We check if the password matches the server pass and if it does:
	 * If the room did not previously exist, we create a new private and public key, encrypt them with the client key
	 * and return it to the original client (so the original client can decrypt them with his own key).
	 * Else, we return the existing ones.
	 */
	rooms[id] = newRoom(user)
	keys[id] = newKeys(key)
	clients[user] = newUser(user, id)
	go roomHandler(id)
}

func addUser (client *websocket.Conn) {
	clients[client] = newUser(client, "")
}