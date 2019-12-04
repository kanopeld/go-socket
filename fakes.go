package socket

import (
	"net"
	"time"
)

//FakeServerClient struct for emulate SClient interface. Used in tests only
type FakeServerClient struct{ Id string }

//ID return the socket id
func (c *FakeServerClient) ID() string { return c.Id }

//Emit sends an event to the other side
func (c *FakeServerClient) Emit(event string, msg interface{}) error { return nil }
func (c *FakeServerClient) send(p *Package) error                    { return nil }

//On registers an event handler under the given name.
func (c *FakeServerClient) On(event string, f interface{}) error { return nil }

//Off deletes an event handler.
func (c *FakeServerClient) Off(event string) bool { return true }

//Disconnect drop current connection. Send the appropriate message to the other side
func (c *FakeServerClient) Disconnect() {}

//Connection return net.Conn with which the socket was created
func (c *FakeServerClient) Connection() net.Conn { return nil }

//Broadcast sends an event to the other side to everyone in the specified room
func (c *FakeServerClient) Broadcast(event string, msg interface{}) error { return nil }

//FakeNetConn the struct for emulate net.Conn interface. Used in tests only
type FakeNetConn struct{}

//Read read doc for net.Conn interface
func (fnc *FakeNetConn) Read(b []byte) (n int, err error) { return 0, nil }

//Write read doc for net.Conn interface
func (fnc *FakeNetConn) Write(b []byte) (n int, err error) { return 0, nil }

//Close read doc for net.Conn interface
func (fnc *FakeNetConn) Close() error { return nil }

//LocalAddr read doc for net.Conn interface
func (fnc *FakeNetConn) LocalAddr() net.Addr { return &net.IPAddr{} }

//RemoteAddr read doc for net.Conn interface
func (fnc *FakeNetConn) RemoteAddr() net.Addr { return &net.IPAddr{} }

//SetDeadline read doc for net.Conn interface
func (fnc *FakeNetConn) SetDeadline(t time.Time) error { return nil }

//SetReadDeadline read doc for net.Conn interface
func (fnc *FakeNetConn) SetReadDeadline(t time.Time) error { return nil }

//SetWriteDeadline read doc for net.Conn interface
func (fnc *FakeNetConn) SetWriteDeadline(t time.Time) error { return nil }
