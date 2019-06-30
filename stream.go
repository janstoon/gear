package gear

import (
	"context"
)

type Stream interface {
	Send(b []byte) error
	Receive() ([]byte, error)
}

type StreamClient interface {
	Stream

	Connect() error
	Shutdown(ctx context.Context) error
}

type StreamGateway interface {
	Stream

	ListenAndServe() error
	ListenAndServeTLS() error
	Shutdown(ctx context.Context) error
}
