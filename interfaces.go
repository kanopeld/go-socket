package socket

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

type handlerSharer interface {
	getEvents() events
	getBroadcast() BroadcastAdaptor
}

type Disconnecter interface {
	Disconnect()
}

//This interface is used as an extension of Emiiter
type IdentifiableEmitter interface {
	IDer
	Emitter
}

//The main server interface. Include Broadcast interface
type SClient interface {
	Client
	Broadcaster
}

//The main client interface. Not include Broadcast interface
type DClient interface {
	Client
}

//Basic client interface
type Client interface {
	Emitter
	IDer
	Connectioner
	Handler
	Disconnecter
}
