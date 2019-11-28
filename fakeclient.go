package socket

import "net"

type FakeClient struct {
	Id string
}

func (c *FakeClient) ID() string {
	return c.Id
}

func (c *FakeClient) Emit(event string, msg interface{}) error {
	return nil
}

func (c *FakeClient) On(event string, f interface{}) error {
	return nil
}

func (c *FakeClient) Off(event string) bool {
	return true
}

func (c *FakeClient) Disconnect() {

}

func (c *FakeClient) Connection() net.Conn {
	return nil
}

func (c *FakeClient) Broadcast(event string, msg []byte) error {
	return nil
}
