package telemetry

func createCon(id int) Connection {
	var connection = DialSNMP("", "", "", "", 0)
	return connection
}

func CreatePool(count int) {
	for i := 0; i < count; i++ {
		go createCon(i)
	}
}
