package socket

import (
	"net"
	"bufio"
	"sync"
)

type dial struct {
	*clientHandler
	*defaultEmitter
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
		defaultEmitter:&defaultEmitter{c:conn},
	}
	d.clientHandler = newClientHandler(d, &baseHandler{events:make(map[string]*caller, 0), name:BASE_HANDLER_DEFAULT_NAME, evMu:sync.Mutex{}})
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
	_, _ = d.conn.Write(p.MarshalBinary())
}

func (d *dial) Disconnect() {
	_ = d.send(&Package{PT:_PACKET_TYPE_DISCONNECT})
	_ = d.call(DISCONNECTION_NAME, nil)
	_ = d.conn.Close()
}

func (d *dial) Broadcast(event string, msg []byte) error {
	return nil
}

func (d *dial) loop() {
	defer d.Disconnect()
	d.sendConnect()

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
				return
			}
		case _PACKET_TYPE_DISCONNECT:
			return
		case _PACKET_TYPE_EVENT:
			msg ,err := DecodeMessage(p.Payload)
			if err != nil {
				continue
			}

			if err := d.call(msg.EventName, msg.Data); err != nil {
				return
			}
		}
	}
}
