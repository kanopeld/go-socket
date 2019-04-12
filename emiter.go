package socket

import "net"

type defaultEmitter struct {
	c net.Conn
}

func (de *defaultEmitter) send(p *Package) error {
	_, err := de.c.Write(p.MarshalBinary())
	return err
}

func (de *defaultEmitter) Emit(event string, data []byte) error {
	return de.send(&Package{PT: _PACKET_TYPE_EVENT, Payload: Message{EventName: event, Data: data}.MarshalBinary()})
}
