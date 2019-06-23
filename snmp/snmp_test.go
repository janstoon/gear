package snmp

import (
	"github.com/k-sone/snmpgo"
	"testing"
)

func TestGetSNMP(t *testing.T) {

	snmpTest := SnmpAgent{
		Username:  "a",
		Password:  "b",
		Community: "public",
		Ip:        "127.0.0.1:1024",
		Version:   snmpgo.V2c,
		Oids:      []string{"1.1.1.1.1.1", "1.3.6.1.2.1.1.9.1.4.8"},
	}
	a, err := snmpTest.GetSNMP()
	if a == nil {
		t.Errorf("Something is wrong amd return value is %v with error: %v", a, err)
	}
}
