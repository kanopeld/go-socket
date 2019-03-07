package socket

import (
	"net"
	"bufio"
	"sync"
)

type dial struct {
	*clientHandler
	conn net.Conn
	id string
}

func NewDial(addr string) (Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	d := &dial{
		conn:conn,
	}
	d.clientHandler = newClientHandler(d, &baseHandler{events:make(map[string]*caller, 0), name:BASE_HANDLER_DEFAULT_NAME, evMu:sync.Mutex{}})
	d.sendConnect()

	go d.loop()
	return d, nil
}

func (d *dial) ID() string {
	return d.id
}

func (d *dial) Connection() net.Conn {
	return d.conn
}

func (d *dial) sendConnect() {
	p := NewPacket(_PACKET_TYPE_CONNECT)
	d.conn.Write(p.MarshalBinary())
}

func (d *dial) sendDisconnect() {
	p := NewPacket(_PACKET_TYPE_DISCONNECT)
	d.conn.Write(p.MarshalBinary())
}

func (d *dial) Disconnect() {
	d.sendDisconnect()
	d.call(DISCONNECTION_NAME, nil)
	d.conn.Close()
}

func (d *dial) send(p *Package) error {
	_, err := d.conn.Write(p.MarshalBinary())
	return err
}

func (d *dial) Emit(event string, data []byte) error {
	b, err := Message{EventName:event,Data:data}.MarshalBinary()
	if err != nil {
		return err
	}
	return d.send(&Package{PT:_PACKET_TYPE_EVENT, Payload:b})
}

func (d *dial) loop() {
	for {
		msg, err := bufio.NewReader(d.conn).ReadBytes('\n')
		if err != nil {
			continue
		}

		p, err := DecodePackage(msg)
		if err != nil {
			continue
		}

		switch p.PT {
		case _PACKET_TYPE_CONNECT:
			d.id = string(p.Payload)
			if err := d.call(CONNECTION_NAME, nil); err != nil {
				d.Disconnect()
				return
			}
		case _PACKET_TYPE_DISCONNECT:
			d.Disconnect()
		case _PACKET_TYPE_EVENT:
			msg ,err := DecodeMessage(p.Payload)
			if err != nil {
				continue
			}

			if err := d.call(msg.EventName, msg.Data); err != nil {
				d.Disconnect()
				return
			}
		}
	}
}
