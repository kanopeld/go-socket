package socket

import "net"

// FakeClient used for testing
type FakeClient struct {
	// Id is the id of the FakeClient
	Id string
}

// ID returns the id of a client
func (c *FakeClient) ID() string {
	return c.Id
}

// Emit sends a message to the server
func (c *FakeClient) Emit(event string, msg interface{}) error {
	return nil
}

// On reacts to incoming messages
func (c *FakeClient) On(event string, f interface{}) error {
	return nil
}

// Off disables a given callback
func (c *FakeClient) Off(event string) bool {
	return true
}

// Disconnect disconnects from the server
func (c *FakeClient) Disconnect() {

}

// Connection returns the connection to the server
func (c *FakeClient) Connection() net.Conn {
	return nil
}

// Broadcast sends a message to everyone in the room
func (c *FakeClient) Broadcast(event string, msg []byte) error {
	return nil
}
