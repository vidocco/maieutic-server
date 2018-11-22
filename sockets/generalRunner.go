package sockets

func ClientHandler () {
	for {
		msg := <- Broadcast

		if msg.Event.Type == "message" {
			// TODO: we need the client here (the user that sent the message...)

		}
	}
}