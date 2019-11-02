package telemetry

import (
	"fmt"
	"gitlab.com/janstun/gear"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Devices map[string]device

type device struct {
	DevName     string
	Protocol    string
	ProtocolOpt ProtocolOptType
	Oids        []string
}

type ProtocolOptType struct {
	Username, Password, Community, Ip, Version string
}

// Construct a new device :
// devName = Device Name eg. dev1
// protocol = snmp , rtu , ...
// protocolOpt >> snmp = username/password/community/ip:port/version
// protocolOpt >> rtu = ...
// protocolOpt >> ... = ...

func (tm Devices) AddDevice(devName, protocol string, protocolOpt ProtocolOptType) device {
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

	pattern := `^\w+\/[1]\.(\w|\.)+\??\d*$`
	validation, _ := regexp.MatchString(pattern, topic)
	if validation == false {
		err := fmt.Errorf("please check your topic: %s again", topic)
		return nil, err
	}

	topicSplited := strings.Split(topic, "/")
	devName := topicSplited[0]
	devParams := strings.Split(topicSplited[1],"?")
	devOid := devParams[0]


	// The interval creation , it is correct but commented because has not been completely coded.
	var devInterval int
	if len(devParams) > 1 {
		devInterval, _ = strconv.Atoi(devParams[1])
		if devInterval == 0 {
			devInterval = 1
		}
	} else {
		devInterval = 1
	}


	var result string
	var err error

	msg := gear.Message{
		Topic: topic,
		Reply: "",
		Data:  []byte(result),
	}

	msgRes := make(chan gear.Message)

	for i, a := range tm {
		if i == devName {
			conn := makeConnection(a)
			for _, o := range a.Oids {
				fmt.Printf("%s , %s , %d , @ %s\n", devParams , devOid , devInterval , o)
				if devOid == o {
					fmt.Println("matched")
					//tick
					fmt.Println(time.Duration(devInterval) * time.Second)
					tick := time.NewTicker(time.Duration(devInterval) * time.Second)
					tickClose := make(chan struct{})
					go func() {
						for {
							select {
							case <-tick.C:
								fmt.Println("tick")
								result, err = conn.Get(o)
								msg.Data = []byte(result)
								msgRes <- msg
								fmt.Printf(" result: %s\n", result)
							case <-tickClose:
								tick.Stop()
								return
							}
						}
					}()
				break
				} else {
					fmt.Println("not matched!")
					return nil , fmt.Errorf("Your desired OID has not been registered: %s" , devOid)
				}
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return msgRes, err
}

func (tm Devices) Unsubscribe(topic string) error {

	return nil
}

// Create a connection to use the desired protocol.
// Now only support snmp protocol
func makeConnection(device device) Connection {

	var connection Connection

	// Make Connection for SNMP :
	if device.Protocol == "snmp" {
		version, _ := strconv.Atoi(device.ProtocolOpt.Version)
		connection = DialSNMP(device.ProtocolOpt.Username, device.ProtocolOpt.Password, device.ProtocolOpt.Community, device.ProtocolOpt.Ip, version)
	}
	return connection
}
