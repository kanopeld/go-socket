package socket

import "net"

type Ider interface {
	// ID returns the socket id
	ID() string
}

type EventHandler interface {
	// On registers an event handler under the given name
	On(event string, c HandlerCallback)
	// Off deletes an event handler. Return true if event existed
	Off(event string) bool
}

// Client includes all basic functions for a client
type Client interface {
	Emitter
	Ider
	// Connection returns net.Conn with which the socket was created
	Connection() net.Conn
	// Disconnect drops current connection. Sends the appropriate message to the other side
	Disconnect()
	EventHandler
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

type broadcastAdapter interface {
	join(room string, c Client) error
	leave(room string, c Client) error
	send(ignore Client, room, event string, msg []byte) error
	len(room string) int
}

//Broadcaster now available only for server-side clients, for dial client it is stub now
type Broadcaster interface {
	//Join add called client into room by given name, if there is no room it will be created
	Join(room string) error
	//Leave remove called client from room by given name, if after deleting the user in the room no one is left she will be deleted too
	Leave(room string) error
	// BroadcastTo sends an event to everyone else in the room by given name
	BroadcastTo(room, event string, msg []byte) error
	// Broadcast sends an event to everyone else in the room default("defaultBroadcast") room
	Broadcast(event string, arg []byte) error
}

type Server interface {
	Start()
	Stop()
	EventHandler
}

// HandlerCallback is function that gets called on a certain event
type HandlerCallback func(c Client, data []byte) error
