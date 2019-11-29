package core

import "net"

type ServerClient interface {
	ID() string
	Connection() net.Conn
	On(event string, f interface{}) error
	Off(event string) bool
	Emit(event string, arg interface{}) error
	Broadcast(event string, arg interface{}) error
	Disconnect()
}

type Client interface {
	ID() string
	Connection() net.Conn
	On(event string, f interface{}) error
	Off(event string) bool
	Emit(event string, arg interface{}) error
	Disconnect()
}

type DialClient interface {
	ID() string
	Connection() net.Conn
	On(event string, f interface{}) error
	Off(event string) bool
	Emit(event string, arg interface{}) error
	Disconnect()
}
