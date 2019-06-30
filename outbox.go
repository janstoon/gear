package gear

type EmailOutbox interface {
	SendEmail(interface{}) error
}

type SmsOutbox interface {
	SendSms(interface{}) error
}
