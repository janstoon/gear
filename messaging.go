package gear

import "time"

type Message struct {
	// Message subject which indicates the queue the message have to enqueue in
	Topic string

	// In case message have to be answered the receiver may use
	// `Reply` as topic to publish the answer
	Reply string

	// Message payload
	Data []byte
}

type Subscriber interface {
	Subscribe(topic string) (<-chan Message, error)
	Unsubscribe(topic string) error
}

type Publisher interface {
	Publish(msg Message) error
}

// Message Queue - Publish Subscribe service
type AsyncBroker interface {
	Publisher
	Subscriber
}

type Applicant interface {
	Request(topic string, data []byte, timeout time.Duration) ([]byte, error)
}

// Message Queue - Request Reply service
type SyncBroker interface {
	Applicant
}

// Message Queue - Push Pull service
type Pipeline interface {
	Push([]byte) error
	Pull() ([]byte, error)
}
