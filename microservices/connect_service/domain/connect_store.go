package domain

type ConnectionStore interface {
	Connect(user, userConnect string) error
	UnConnect() error
}
