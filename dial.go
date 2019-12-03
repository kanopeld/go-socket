package socket

import (
	"bufio"

	"net"
)

type dial struct {
	*dialHandler
	Emitter
	conn net.Conn
	id   string
	disc bool
}

func NewDial(addr string) (DClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	d := &dial{
		conn:    conn,
		Emitter: GetEmitter(conn),
	}
	d.dialHandler = newDialHandler(d)
	go d.loop()
	return d, nil
}

func (d *dial) ID() string {
	return d.id
}

func (d *dial) Connection() net.Conn {
	return d.conn
}

func (d *dial) Disconnect() {
	if d.disc {
		return
	}
	d.disc = true
	_ = d.Send(&Package{PT: PackTypeDisconnect})
	_ = d.call("disconnection", nil)
	_ = d.conn.Close()
}

func (d *dial) loop() {
	defer func() {
		d.Disconnect()
	}()

	reader := bufio.NewReader(d.conn)
	for {
		msg, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		p, err := decodePackage(msg)
		if err != nil {
			return
		}

		switch p.PT {
		case PackTypeConnect:
			d.id = string(p.Payload)
			if err = d.Send(&Package{PT: PackTypeConnectAccept}); err != nil {
				return
			}

			if err := d.call("connection", nil); err != nil {
				return
			}
		case PackTypeDisconnect:
			return
		case PackTypeEvent:
			msg := decodeMessage(p.Payload)
			if err := d.call(msg.EventName, msg.Data); err != nil {
				return
			}
		}
	}
}
