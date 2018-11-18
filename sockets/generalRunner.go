package sockets

func ClientHandler () {
	for {
		msg := <- Broadcast

		if msg.Type == "message" {

		}
	}
}