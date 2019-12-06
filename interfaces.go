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
	On(event string, c HandlerCallback)
	//Off deletes an event handler. Return true if event was exist
	Off(event string) bool
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

//Basic client interface. Includes all basic functions
type Client interface {
	Emitter
	IDer
	Connectioner
	Handler
	Disconnecter
	Broadcaster
}

type looper interface {
	loop()
}

type HandlerCallback func(c Client, data []byte) error
