package socket

import (
	"net"
	"time"
)

type FakeServerClient struct{ Id string }

func (c *FakeServerClient) ID() string                                    { return c.Id }
func (c *FakeServerClient) Emit(event string, msg interface{}) error      { return nil }
func (c *FakeServerClient) Send(p *Package) error                         { return nil }
func (c *FakeServerClient) On(event string, f interface{}) error          { return nil }
func (c *FakeServerClient) Off(event string) bool                         { return true }
func (c *FakeServerClient) Disconnect()                                   {}
func (c *FakeServerClient) Connection() net.Conn                          { return nil }
func (c *FakeServerClient) Broadcast(event string, msg interface{}) error { return nil }

type FakeNetComm struct{}

func (fnc *FakeNetComm) Read(b []byte) (n int, err error)   { return 0, nil }
func (fnc *FakeNetComm) Write(b []byte) (n int, err error)  { return 0, nil }
func (fnc *FakeNetComm) Close() error                       { return nil }
func (fnc *FakeNetComm) LocalAddr() net.Addr                { return &net.IPAddr{} }
func (fnc *FakeNetComm) RemoteAddr() net.Addr               { return &net.IPAddr{} }
func (fnc *FakeNetComm) SetDeadline(t time.Time) error      { return nil }
func (fnc *FakeNetComm) SetReadDeadline(t time.Time) error  { return nil }
func (fnc *FakeNetComm) SetWriteDeadline(t time.Time) error { return nil }
