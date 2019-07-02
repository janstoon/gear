package telemetry

import (
	"fmt"
	"gitlab.com/janstun/gear"
	"strconv"
	"strings"
)

type TelemetryMessage struct {
}

type device struct {
	DevName     string
	Protocol    string
	ProtocolOpt string
	Oids        []string
}

//To keep Devices in Memory
var tempDevice []device

//Construct a new device :
// devName = Device Name eg. dev1
// protocol = snmp , rtu , ...
// protocolOpt >> snmp = username/password/community/ip:port/version
// protocolOpt >> rtu = ...
// protocolOpt >> ... = ...
func (tm TelemetryMessage) AddDevice(devName, protocol, protocolOpt string) device {

	dev := device{}
	dev.DevName = devName
	dev.Protocol = protocol
	dev.ProtocolOpt = protocolOpt

	if len(tempDevice) > 0 {
		for _, a := range tempDevice {
			aname := a.DevName
			if aname == dev.DevName {
				fmt.Println("There is a duplicate")
			} else {
				tempDevice = append(tempDevice, dev)
			}
		}
	} else {
		tempDevice = append(tempDevice, dev)
	}

	return dev
}

//Set Which Parameter of the device needs to read/write
func (tm TelemetryMessage) WatchParam(devName string, oids []string) {
	var devIterator int

	for i, a := range tempDevice {
		if devName == a.DevName {
			devIterator = i
		}
	}
	tempDevice[devIterator].Oids = oids
}

//Create a Subscription upon on the name of topic : <device name>/<optional description>
func (tm TelemetryMessage) Subscribe(topic string) (<-chan gear.Message, error) {

	devName := strings.Split(topic, "/")

	var devIterator int

	for i, a := range tempDevice {
		if devName[0] == a.DevName {
			devIterator = i
		}
	}

	result, err := makeConnection(tempDevice[devIterator])
	if err != nil {
		return nil, err
	}

	msg := gear.Message{
		Topic: topic,
		Reply: result,
		Data:  nil,
	}

	msgRes := make(chan gear.Message)

	go func() {
		msgRes <- msg
		close(msgRes)
	}()

	return msgRes, err
}

func (tm TelemetryMessage) Unsubscribe(topic string) error {

	return nil
}

// Create a connection to use the desired protocol.
// Now only supported snmp for one oid!
func makeConnection(device device) (string, error) { //func makeConnection(device device) ([]string, error)
	var result string
	var err error
	protocolOpts := strings.Split(device.ProtocolOpt, "/")

	if device.Protocol == "snmp" {
		version, _ := strconv.Atoi(protocolOpts[4])
		connection := DialSNMP(protocolOpts[0], protocolOpts[1], protocolOpts[2], protocolOpts[3], version)
		result, err = connection.Get(device.Oids[0])

		//if Len(device.Oids) > 1 {
		//	result, err := connection.GetMany(device.Oids)
		//} else {
		//  result, err = connection.Get(device.Oids[0])
		//}

	}
	return result, err
}
