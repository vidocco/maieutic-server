package handlers

import (
	"github.com/gorilla/websocket"
	"github.com/naoina/denco"
	"log"
	"maieutic-server/env"
	"maieutic-server/sockets"
	"maieutic-server/utils"
	"net/http"
	"strconv"
)

var rBuffSize, _ = strconv.Atoi(env.GetOr("R_BUFF_SIZE", "1024"))
var wBuffSize, _ = strconv.Atoi(env.GetOr("W_BUFF_SIZE", "1024"))

func origCheck (r *http.Request) bool {
	if env.GetOr("GO_ENV", "development") == "development" {
		return true
	}
	return false
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  rBuffSize,
	WriteBufferSize: wBuffSize,
	CheckOrigin: origCheck,
}

func WsHandler (w http.ResponseWriter, r *http.Request, params denco.Params) {
	conn, err := upgrader.Upgrade(w, r, nil)
	utils.CheckErr(err)

	defer conn.Close()
	for {
		var msg sockets.Event
		err := conn.ReadJSON(&msg)

		if err != nil {
			log.Printf("error: %v", err)
			sockets.Disconnect(conn)
			break
		}

		sockets.Broadcast <- msg
	}
}
