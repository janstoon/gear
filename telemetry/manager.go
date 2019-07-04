package telemetry

import (
	"fmt"
	"gitlab.com/janstun/gear"
	"regexp"
	"strconv"
	"strings"
)

type Devices map[string]device

type device struct {
	DevName     string
	Protocol    string
	ProtocolOpt string
	Oids        []string
}

// Construct a new device :
// devName = Device Name eg. dev1
// protocol = snmp , rtu , ...
// protocolOpt >> snmp = username/password/community/ip:port/version
// protocolOpt >> rtu = ...
// protocolOpt >> ... = ...

func (tm Devices) AddDevice(devName, protocol, protocolOpt string) device {
	dev := device{}
	dev.DevName = devName
	dev.Protocol = protocol
	dev.ProtocolOpt = protocolOpt
	tm[devName] = dev
	return dev
}

// Register Which Parameter of the device needs to read/write
func (tm Devices) RegisterDeviceParams(devName string, oids []string) {
	for i, a := range tm {
		if devName == i {
			a.Oids = append(a.Oids, oids...)
			tm[i] = a
		}
	}
}

// Create a Subscription upon on the name of topic : <device name>/<oid>/<optional interval in seconds>
// e.g : dev1/1.2.1.4.5.1.2/100 or testdev/1.2.1.4.2.1
func (tm Devices) Subscribe(topic string) (<-chan gear.Message, error) {

	pattern := `^\w+\/[1]\.(\w|\.)+\/?\d*$`
	validation, _ := regexp.MatchString(pattern, topic)
	if validation == false {
		err := fmt.Errorf("please check your topic: %s again", topic)
		return nil, err
	}

	topicSplited := strings.Split(topic, "/")
	devName := topicSplited[0]
	devOid := topicSplited[1]

	/* The interval creation , it is correct but commented because has not been completely coded.
	var devInterval int

	if len(topicSplited) > 2 {
		devInterval, _ = strconv.Atoi(topicSplited[2])
	} else {
		devInterval = 1
	}
	*/

	var result string
	var err error

	//tick := time.NewTicker(time.Second * time.Duration(devInterval))
	//go func() {
	//	for {
	//		select {
	//		case <- tick.C:
	for i, a := range tm {
		if i == devName {
			for _, o := range a.Oids {
				if o == devOid {
					result, err = getData(a, o)
				}
			}
		}
	}
	//		}
	//	}
	//}()

	if err != nil {
		return nil, err
	}

	msg := gear.Message{
		Topic: topic,
		Reply: nil,
		Data:  []byte(result),
	}

	msgRes := make(chan gear.Message)

	msgRes <- msg
	close(msgRes)
	return msgRes, err
}

func (tm Devices) Unsubscribe(topic string) error {

	return nil
}

// Create a connection to use the desired protocol.
// Now only supported snmp for one oid!
func getData(device device, oid string) (string, error) { //func makeConnection(device device) ([]string, error)

	var result string
	var err error

	protocolOpts := strings.Split(device.ProtocolOpt, "/")
	fmt.Println(device.Oids)
	if device.Protocol == "snmp" {
		version, _ := strconv.Atoi(protocolOpts[4])
		connection := DialSNMP(protocolOpts[0], protocolOpts[1], protocolOpts[2], protocolOpts[3], version)
		result, err = connection.Get(oid)
	}

	return result, err
}
