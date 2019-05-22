package http

import (
	"context"
	"crypto/tls"
	"net/http"

	"gitlab.com/janstun/actor"
)

type gateway struct {
	mux    *http.ServeMux
	server *http.Server
}

func (gw gateway) Handle(pattern string, handler http.Handler) {
	gw.mux.Handle(pattern, handler)
}

func (gw gateway) ListenAndServe() error {
	return gw.server.ListenAndServe()
}

func (gw gateway) ListenAndServeTLS() error {
	return gw.server.ListenAndServeTLS("", "")
}

func (gw gateway) Shutdown(ctx context.Context) error {
	return gw.server.Shutdown(ctx)
}

func NewGateway(addr string, tlsConfig *tls.Config) actor.HttpGateway {
	gw := &gateway{}
	gw.mux = http.NewServeMux()
	gw.server = &http.Server{
		Addr:      addr,
		Handler:   gw.mux,
		TLSConfig: tlsConfig,
	}

	return gw
}
