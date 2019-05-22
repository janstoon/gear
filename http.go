package actor

import (
	"context"
	"net/http"
)

type HttpMux interface {
	Handle(pattern string, handler http.Handler)
}

type mmux struct {
	muxes []HttpMux
}

type HttpGateway interface {
	HttpMux

	ListenAndServe() error
	ListenAndServeTLS() error
	Shutdown(ctx context.Context) error
}

func CombineHttpMuxes(muxes ...HttpMux) HttpMux {
	return mmux{muxes: muxes}
}

func (m mmux) Handle(pattern string, handler http.Handler) {
	for _, mux := range m.muxes {
		mux.Handle(pattern, handler)
	}
}
