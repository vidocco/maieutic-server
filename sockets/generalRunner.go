package sockets

func clientHandler () {
	for {
		msg := <- Broadcast

		if msg.Event.Type == "init" {
			addUser(msg.Client)
		} else if msg.Event.Type == "message" {
			rooms[clients[msg.Client].Room].Broadcast <- msg.Event
		} else if msg.Event.Type == "create room" {

		} else if msg.Event.Type == "join room" {
			
		}
	}
}

func SpawnSocketHandlers (n int) {
	for i := 0; i < n; i++ {
		go clientHandler()
	}
}