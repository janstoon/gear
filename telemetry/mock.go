package telemetry

type mockConn int

func (m mockConn) Get(id string) (string, error) {
	return "ok", nil
}

func (m mockConn) GetMany(id []string) (map[string]string, error) {

	mapres := make(map[string]string)
	return mapres, nil

}

func (m mockConn) Set(id string, value string) error {

	return nil
}
