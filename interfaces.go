package socket

import "net"

type Ider interface {
	// ID returns the socket id
	ID() string
}

// Client includes all basic functions for a client
type Client interface {
	Emitter
	Ider
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
	SetClient(c Client) error
	RemoveClient(c Client) error
	Len() int
	Send(ignore Client, event string, msg []byte) error
	ClientExists(c Client) (ok bool)
}

type Broadcaster interface {
	Join(room string, c Client) error
	Leave(room string, c Client) error
	Send(ignore Client, room, event string, msg []byte) error
	Len(room string) int
}

// HandlerCallback is function that gets called on a certain event
type HandlerCallback func(c Client, data []byte) error
