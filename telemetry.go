package actor

type Telemetry interface {
	Get(id string) (string, error)
	GetMany(id []string) (map[string]string, error)
	Set(id string, value string) error
}
