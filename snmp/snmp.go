package snmp

import (
	"fmt"
	"github.com/k-sone/snmpgo"
)

type SnmpAgent struct {
	Username  string
	Password  string
	Community string
	Ip        string
	Version   snmpgo.SNMPVersion
	Oids      []string
}

func (sa SnmpAgent) GetSNMP() (snmpgo.Variable, error) {

	//Creating SNMP Object :
	snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   sa.Version, //SNMP Version eg. v1 or v3
		Address:   sa.Ip,      //Device Address
		Retries:   1,
		Community: sa.Community, //Community string that defines in the device menu. For v3 maybe contains user and password.
	})

	if err != nil {
		// Failed to create snmpgo.SNMP object
		return nil, err
	}

	//Oids :
	//oids, err := snmpgo.NewOids([]string{})

	/* Config Oids : */
	oids, err := snmpgo.NewOids([]string{
		sa.Oids[0],
		sa.Oids[1],
	})

	//Open Connection:
	if err = snmp.Open(); err != nil {
		// Failed to open connection
		return nil, err
	}
	defer snmp.Close()

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

	// get VarBind list : pdu.VarBinds())

	// select a VarBind : pdu.VarBinds().MatchOid(oids[0])

	return pdu.VarBinds().MatchOid(oids[1]).Variable, err
}
