package core

import "net"

type IDer interface {
	ID() string
}

type Connectioner interface {
	Connection() net.Conn
}

type Handler interface {
	On(event string, f interface{}) error
	Off(event string) bool
}

type HandlerSharer interface {
	GetEvents() Events
	GetBroadcast() BroadcastAdaptor
}

type Disconnecter interface {
	Disconnect()
}

type IdentifiableEmitter interface {
	IDer
	Emitter
}

type SClient interface {
	Client
	Broadcaster
}

type DClient interface {
	Client
}

type Client interface {
	Emitter
	IDer
	Connectioner
	Handler
	Disconnecter
}
