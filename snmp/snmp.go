package snmp

import (
	"fmt"
	"github.com/k-sone/snmpgo"
	"net"
)

type snmpAgent struct {
	username  string
	password  string
	community string
	ip        net.IPAddr
	version   snmpgo.SNMPVersion
	oids      []string
}

func (sa snmpAgent) getSNMP() (string, error) {

	//Creating SNMP Object :
	snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   sa.version,     //SNMP Version eg. v1 or v3
		Address:   sa.ip.String(), //Device Address
		Retries:   1,
		Community: sa.community, //Community string that defines in the device menu. For v3 maybe contains user and password.
	})

	if err != nil {
		// Failed to create snmpgo.SNMP object
		fmt.Println(err)
		return "", err
	}

	//Oids :
	oids, err := snmpgo.NewOids([]string{})

	/* Config Oids :
	oids, err := snmpgo.NewOids([]string{
		"1.3.6.1.2.1.1.1.0",
		"1.3.6.1.2.1.1.2.0",
		"1.3.6.1.2.1.1.3.0",
	})

	*/

	//Open Connection:
	if err = snmp.Open(); err != nil {
		// Failed to open connection
		fmt.Println(err)
		return "", err
	}
	defer snmp.Close()

	//GetData :
	pdu, err := snmp.GetRequest(oids)
	if err != nil {
		// Failed to request
		fmt.Println(err)
		return "", err
	}
	if pdu.ErrorStatus() != snmpgo.NoError {
		// Received an error from the agent
		fmt.Println(pdu.ErrorStatus(), pdu.ErrorIndex())
	}

	// get VarBind list
	fmt.Println(pdu.VarBinds())

	// select a VarBind
	fmt.Println(pdu.VarBinds().MatchOid(oids[0]))

	return pdu.VarBinds().String(), nil
}
