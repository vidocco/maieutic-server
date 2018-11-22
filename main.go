package main

import (
	"log"
	"maieutic-server/env"
	"maieutic-server/middlewares"
	"maieutic-server/router"
	"maieutic-server/sockets"
	"maieutic-server/utils"
	"net/http"
	"strconv"
)

func main() {
	r := router.NewRouter()
	handler := middlewares.ApplyMiddleware(r)
	log.Printf("Frank running in %s%s", env.GetOr("MAIEUTIC_HOST", "localhost"), env.GetOr("MAIEUTIC_PORT", ":470"))
	err := http.ListenAndServe(env.GetOr("MAIEUTIC_PORT", ":470"), handler)
	utils.CheckErr(err)
}

func init() {
	env.SetEnv()
	number, err := strconv.Atoi(env.GetOr("SOCKET_HANDLERS", "4"))
	utils.CheckErr(err)
	sockets.SpawnSocketHandlers(number)
}