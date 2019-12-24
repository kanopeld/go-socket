package socket

import (
	"net"
)

// Emitter organizes sending events to the other side
type Emitter interface {
	// Emit sends an event to the other side
	Emit(event string, arg []byte) error
	sender
}

func getEmitter(c net.Conn) Emitter {
	return &defaultEmitter{getSender(c)}
}

type defaultEmitter struct {
	sender
}

// Emit sends an event to the other side
func (de *defaultEmitter) Emit(event string, arg []byte) error {
	return de.send(&sockPackage{PT: PackTypeEvent, Payload: message{EventName: event, Data: arg}.MarshalBinary()})
}
