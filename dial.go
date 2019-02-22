package socket

import (
	"net"
	"bufio"
)

type Dial struct {
	c net.Conn
	ad ClientAdaptor
	l bool
	cl bool
}

func NewDial(addr string) (*Dial, error) {
	d := &Dial{
		ad:GetClientAdapptor(),
	}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	d.c = conn

	d.sendConnect()

	go d.loop()
	return d, nil
}

func (d *Dial) sendConnect() {
	p := NewPacket(_PACKET_TYPE_CONNECT)
	d.c.Write(p.MarshalBinary())
}

func (d *Dial) sendDisconnect() {
	p := NewPacket(_PACKET_TYPE_DISCONNECT)
	d.c.Write(p.MarshalBinary())
}

func (d *Dial) sendError() {
	p := NewPacket(_PACKET_TYPE_ERROR)
	d.c.Write(p.MarshalBinary())
}

func (d *Dial) send(t PackageType, name string, data string) {
	msg := Message{
		EventName:name,
		Data:MessagePayload{Data:[]byte(data)},
	}

	p, err := NewEventPacket(msg)
	if err != nil {
		return
	}

	d.c.Write(p.MarshalBinary())
}

func (d *Dial) loop() {
	defer func() {
		d.Close()
	}()

	if d.l || d.cl {
		return
	}

	d.l = true
	for d.l {
		msg, err := bufio.NewReader(d.c).ReadBytes('\n')
		if err != nil {
			continue
		}

		p, err := DecodePackage(msg)
		if err != nil {
			continue
		}

		switch p.PT {
		case _PACKET_TYPE_DISCONNECT:
			d.ad.Call(DISCONNECTION_NAME, "")
			d.Close()
		case _PACKET_TYPE_EVENT:
			msg ,err := DecodeMessage(p.Payload)
			if err != nil {
				continue
			}

			d.ad.Call(msg.EventName, msg.Data.String())
		default:
			d.ad.Call(ERROR_EVENT, "")
			d.ad.Call(DISCONNECTION_NAME, "")
			d.Close()
		}
	}
}

func (d *Dial) Close() {
	d.sendDisconnect()
	d.cl = true
	d.l = false
	d.c.Close()
}

func (d *Dial) On(name string, callback ClientEventCallback) {
	d.ad.On(name, callback)
}

func (d *Dial) Send(name string, data string) {
	d.send(_PACKET_TYPE_EVENT, name, data)
}