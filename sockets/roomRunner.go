package sockets

func roomHandler (id string) {
	for {
		msg := <- rooms[id].Broadcast

		if msg.Type == "disconnect" {
			break
		} else if msg.Type == "message" {
			for _, user := range rooms[id].Users {
				user.WriteJSON(msg.Content)
			}
		}
	}
}