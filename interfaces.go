package socket

import "net"

type ider interface {
	// ID returns the socket id
	ID() string
}

// identifiableEmitter is an extension of Emiter
type identifiableEmitter interface {
	ider
	// Emitter organizes sending events to the other side
	Emitter
}

// Client includes all basic functions for a client
type Client interface {
	Emitter
	ider
	// Broadcast sends an event to everyone else in the room
	Broadcast(event string, arg []byte) error
	// Connection returns net.Conn with which the socket was created
	Connection() net.Conn
	// Disconnect drops current connection. Sends the appropriate message to the other side
	Disconnect()
	// On registers an event handler under the given name
	On(event string, c HandlerCallback)
	// Off deletes an event handler. Return true if event existed
	Off(event string) bool
}

type looper interface {
	loop()
}

type roomer interface {
	SetClient(c identifiableEmitter) error
	RemoveClient(c identifiableEmitter) error
	Len() int
	Send(ignore identifiableEmitter, event string, msg []byte) error
	ClientExists(c identifiableEmitter) (ok bool)
}

type broadcaster interface {
	Join(room string, c identifiableEmitter) error
	Leave(room string, c identifiableEmitter) error
	Send(ignore identifiableEmitter, room, event string, msg []byte) error
	Len(room string) int
}

// HandlerCallback is function that gets called on a certain event
type HandlerCallback func(c Client, data []byte) error
