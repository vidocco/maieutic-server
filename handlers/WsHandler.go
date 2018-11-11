package handlers

import (
	"github.com/naoina/denco"
	"net/http"
)

type msg struct {
	Type string
}

func WsHandler (w http.ResponseWriter, r *http.Request, params denco.Params) {
	w.WriteHeader(http.StatusOK)
}
