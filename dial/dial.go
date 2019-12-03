package dial

import (
	"bufio"
	"github.com/kanopeld/go-socket/core"
	"net"
)

type dial struct {
	*clientHandler
	core.Emitter
	conn net.Conn
	id   string
	disc bool
}

func NewDial(addr string) (core.DClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	d := &dial{
		conn:    conn,
		Emitter: core.GetEmitter(conn),
	}
	d.clientHandler = newClientHandler(d)
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
	_ = d.Send(&core.Package{PT: core.PackTypeDisconnect})
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

		p, err := core.DecodePackage(msg)
		if err != nil {
			return
		}

		switch p.PT {
		case core.PackTypeConnect:
			d.id = string(p.Payload)
			if err = d.Send(&core.Package{PT: core.PackTypeConnectAccept}); err != nil {
				return
			}

			if err := d.call("connection", nil); err != nil {
				return
			}
		case core.PackTypeDisconnect:
			return
		case core.PackTypeEvent:
			msg := core.DecodeMessage(p.Payload)
			if err := d.call(msg.EventName, msg.Data); err != nil {
				return
			}
		}
	}
}
