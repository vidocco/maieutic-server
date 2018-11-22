package sockets

func clientHandler () {
	for {
		msg := <- Broadcast

		if msg.Event.Type == "message" {
			// TODO: we need the client here (the user that sent the message...)

		}
	}
}

func SpawnSocketHandlers (n int) {
	for i := 0; i < n; i++ {
		go clientHandler()
	}
}