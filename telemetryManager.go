package gear

type TelemetryMessage struct {
	Topic string
	Reply string
	Data  []byte
}

type TelemetrySubscriber interface {
	Subscribe(topic string) (<-chan Message, error)
	Unsubscribe(topic string) error
}

type device struct{
	DevName string
	Protocol string
	ProtocolOpt string
	Oids []string
}

func (tm TelemetryMessage)AddDevice(devName , protocol , protocolOpt string)device {
	dev := device{}
	dev.DevName = devName
	dev.Protocol = protocol
	dev.ProtocolOpt = protocolOpt
	return dev
}

func (dev device)WatchParam(oids []string) {
	dev.Oids = oids
}

func (tm TelemetryMessage)Subscribe(topic string)(<-chan Message , error){

	return nil , nil
}
