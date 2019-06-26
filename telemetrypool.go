package actor

import "gitlab.com/janstun/actor/telemetry"

func createCon(id int) telemetry.Connection {
	var connection = telemetry.DialSNMP("", "", "", "", 0)
	return connection
}

func CreatePool(count int) {
	for i := 0; i < count; i++ {
		go createCon(i)
	}
}
