package actor

type BasicAuthentication interface {
	BasicAuthenticate(username, password string) error
}

type KeyAuthentication interface {
	KeyAuthenticate(key string) error
}
