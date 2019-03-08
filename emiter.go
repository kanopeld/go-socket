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
	b, err := Message{EventName:event,Data:data}.MarshalBinary()
	if err != nil {
		return err
	}
	return de.send(&Package{PT:_PACKET_TYPE_EVENT, Payload:b})
}
