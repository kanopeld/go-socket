package socket

const (
	// ConnectionName is the name of an event that is called when new client has connected (server side) or when client connected to the server (client side)
	ConnectionName = "connection"
	// DisconnectionName is the name of an event that is called when client has disconnected
	DisconnectionName = "disconnection"
)
