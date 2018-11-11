package router

import (
	"github.com/naoina/denco"
	"maieutic-server/handlers"
	"maieutic-server/utils"
	"net/http"
)

func NewRouter() http.Handler {
	mux := denco.NewMux()

	router, err := mux.Build([]denco.Handler{
		mux.GET("/", handlers.Base),
		mux.GET("/ws", handlers.WsHandler),
	})

	utils.CheckErr(err)
	return router
}