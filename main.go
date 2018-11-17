package main

import (
	"log"
	"maieutic-server/env"
	"maieutic-server/middlewares"
	"maieutic-server/router"
	"maieutic-server/utils"
	"net/http"
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
}