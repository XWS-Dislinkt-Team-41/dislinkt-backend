package domain

type ConnectStore interface {
	Connect(user, userConnect string) error
	UnConnect(user, userConnect string) error
	GetUserConnections(user string) ([]string, error)
}

type ConnectRequestStore interface {
}
