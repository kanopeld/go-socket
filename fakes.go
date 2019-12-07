package socket

import (
	"net"
	"time"
)

// fakeServerClient struct for emulate SClient interface. Used in tests only
type fakeServerClient struct{ id string }

// ID return the socket id
func (c *fakeServerClient) ID() string { return c.id }

// Emit sends an event to the other side
func (c *fakeServerClient) Emit(event string, msg []byte) error { return nil }
func (c *fakeServerClient) send(p *Package) error               { return nil }

// On registers an event handler under the given name.
func (c *fakeServerClient) On(event string, f HandlerCallback) {}

// Off deletes an event handler. Return true if event was exist
func (c *fakeServerClient) Off(event string) bool { return true }

// Disconnect drop current connection. Send the appropriate message to the other side
func (c *fakeServerClient) Disconnect() {}

// Connection return net.Conn with which the socket was created
func (c *fakeServerClient) Connection() net.Conn { return nil }

// Broadcast sends an event to the other side to everyone in the specified room
func (c *fakeServerClient) Broadcast(event string, msg []byte) error { return nil }

// fakeNetConn the struct for emulate net.Conn interface. Used in tests only
type fakeNetConn struct{}

// Read read doc for net.Conn interface
func (fnc *fakeNetConn) Read(b []byte) (n int, err error) { return 0, nil }

// Write read doc for net.Conn interface
func (fnc *fakeNetConn) Write(b []byte) (n int, err error) { return 0, nil }

// Close read doc for net.Conn interface
func (fnc *fakeNetConn) Close() error { return nil }

// LocalAddr read doc for net.Conn interface
func (fnc *fakeNetConn) LocalAddr() net.Addr { return &net.IPAddr{} }

// RemoteAddr read doc for net.Conn interface
func (fnc *fakeNetConn) RemoteAddr() net.Addr { return &net.IPAddr{} }

// SetDeadline read doc for net.Conn interface
func (fnc *fakeNetConn) SetDeadline(t time.Time) error { return nil }

// SetReadDeadline read doc for net.Conn interface
func (fnc *fakeNetConn) SetReadDeadline(t time.Time) error { return nil }

// SetWriteDeadline read doc for net.Conn interface
func (fnc *fakeNetConn) SetWriteDeadline(t time.Time) error { return nil }
