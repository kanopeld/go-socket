package socket

import "net"

type IDer interface {
	//ID return the socket id
	ID() string
}

type Connectioner interface {
	//Connection return net.Conn with which the socket was created
	Connection() net.Conn
}

type Handler interface {
	//On registers an event handler under the given name.
	On(event string, f interface{}) error
	//Off deletes an event handler.
	Off(event string) bool
}

type handlerSharer interface {
	getEvents() events
	getBroadcast() BroadcastAdaptor
}

type Disconnecter interface {
	//Disconnect drop current connection. Send the appropriate message to the other side
	Disconnect()
}

//This interface is used as an extension of Emiiter
type IdentifiableEmitter interface {
	IDer
	//Emitter organizes sending events to the other side
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
	//Emitter organizes sending events to the other side
	Emitter
	IDer
	Connectioner
	Handler
	Disconnecter
}

type looper interface {
	loop()
}
