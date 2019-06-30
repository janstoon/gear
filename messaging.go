package gear

import "time"

type Message struct {
	Topic string
	Reply string
	Data  []byte
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
