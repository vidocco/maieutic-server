package middlewares

import "net/http"

type Middleware func (handler http.Handler) http.Handler

type Middlewares []Middleware

var stack = Middlewares{
	logger,
	cors,
}