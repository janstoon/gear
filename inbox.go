package gear

type EmailInbox interface {
	ReceiveEmail() (interface{}, error)
}

type SmsInbox interface {
	ReceiveSms() (interface{}, error)
}
