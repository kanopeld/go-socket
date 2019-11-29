package core

import "net"

type FakeServerClient struct{ Id string }

func (c *FakeServerClient) ID() string                               { return c.Id }
func (c *FakeServerClient) Emit(event string, msg interface{}) error { return nil }
func (c *FakeServerClient) On(event string, f interface{}) error     { return nil }
func (c *FakeServerClient) Off(event string) bool                    { return true }
func (c *FakeServerClient) Disconnect()                              {}
func (c *FakeServerClient) Connection() net.Conn                     { return nil }
func (c *FakeServerClient) Broadcast(event string, msg []byte) error { return nil }
