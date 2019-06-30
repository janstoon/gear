package ldap

import (
	"fmt"

	"github.com/janstoon/ldap"
	"gitlab.com/janstun/gear"
)

type basic struct {
	conn   *ldap.Conn
	baseDn string
}

func (s basic) BasicAuthenticate(username, password string) error {
	_, err := s.conn.DigestMd5Bind(ldap.NewDigestMd5BindRequest(fmt.Sprintf("cn=%s,%s", username, s.baseDn), password, "", nil))

	return err
}

func NewBasicAuthentication(network, addr, baseDn string) (gear.BasicAuthentication, error) {
	if conn, err := ldap.Dial(network, addr); err != nil {
		return nil, err
	} else {
		return &basic{conn: conn, baseDn: baseDn}, nil
	}
}
