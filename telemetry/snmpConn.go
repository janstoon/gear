package telemetry

import (
	"fmt"
	"github.com/k-sone/snmpgo"
)

//Struct needed for creating the first snmp objects
type snmpConn struct {
	username  string
	password  string
	community string
	ip        string
	version   int
}

func DialSNMP(username, password, community, ip string, version int) Connection {
	snmpConn := new(snmpConn)
	snmpConn.username = username
	snmpConn.password = password
	snmpConn.community = community
	snmpConn.ip = ip
	snmpConn.version = version
	return snmpConn
}

func (s snmpConn) Get(id string) (string, error) {
	snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.SNMPVersion(s.version), //SNMP Version eg. v1 or v3
		Address:   s.ip,                          //Device Address
		Retries:   1,
		Community: s.community, //Community string that defines in the device menu. For v3 maybe contains user and password.
	})
	if err != nil {
		// Failed to create snmpgo.SNMP object
		return "", err
	}

	if err = snmp.Open(); err != nil {
		// Failed to open connection
		return "", err
	}
	defer snmp.Close()

	idoid := []string{id}
	oids, err := snmpgo.NewOids(idoid)

	//GetData :
	pdu, err := snmp.GetRequest(oids)

	if err != nil {
		// Failed to request
		return "", err
	}

	if pdu.ErrorStatus() != snmpgo.NoError {
		// Received an error from the agent
		fmt.Println(pdu.ErrorStatus(), pdu.ErrorIndex())
	}

	return pdu.VarBinds().MatchOid(oids[0]).Variable.String(), err
}

func (s snmpConn) GetMany(id []string) (map[string]string, error) {
	snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.SNMPVersion(s.version), //SNMP Version eg. v1 or v3
		Address:   s.ip,                          //Device Address
		Retries:   1,
		Community: s.community, //Community string that defines in the device menu. For v3 maybe contains user and password.
	})
	if err != nil {
		// Failed to create snmpgo.SNMP object
		return nil, err
	}

	if err = snmp.Open(); err != nil {
		// Failed to open connection
		return nil, err
	}
	defer snmp.Close()

	oids, err := snmpgo.NewOids(id)

	//GetData :
	pdu, err := snmp.GetRequest(oids)
	if err != nil {
		// Failed to request
		return nil, err
	}
	if pdu.ErrorStatus() != snmpgo.NoError {
		// Received an error from the agent
		fmt.Println(pdu.ErrorStatus(), pdu.ErrorIndex())
	}
	pdu.VarBinds().MatchOid(oids[0]).Variable.String()
	resmap := make(map[string]string)

	for i, k := range id {
		resmap[k] = pdu.VarBinds().MatchOid(oids[i]).Variable.String()
		i++
	}

	return resmap, err
}

func (s snmpConn) Set(id string, value string) error {
	return nil
}
