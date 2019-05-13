package websocket

import (
	"net/http"

	"gitlab.com/janstun/actor"
	"golang.org/x/net/websocket"
)

// WebSocket Http Mux

type mux struct {
	addr string
}

func (s mux) Handle(pattern string, handler http.Handler) {
	http.Handle(pattern, s.createHttpHandler(handler))
}

func (s mux) Serve() error {
	return http.ListenAndServe(s.addr, nil)
}

func (s mux) createHttpHandler(handler http.Handler) websocket.Handler {
	return websocket.Handler(func(conn *websocket.Conn) {
		//TODO: Translate from http handler into web-socket handler
	})
}

func NewHttpMux(addr string) (actor.HttpMux, error) {
	return mux{addr}, nil
}
