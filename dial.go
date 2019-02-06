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

	d.send(_CONNECTION, CONNECTION_NAME, "")

	go d.loop()
	return d, nil
}

func (d *Dial) send(t int, name string, data string) error {
	msg, err := NewEncodeMessage(t, name, data)
	if err != nil {
		return err
	}

	_, err = d.c.Write([]byte(msg))
	if err != nil {
		return err
	}

	return nil
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

		dec, err := NewMessageDecoder(msg)
		if err != nil {
			continue
		}

		switch dec.mt {
		case _EVENT:
			d.ad.Call(dec.eventName, dec.payload)
		}
	}
}

func (d *Dial) Close() {
	d.send(_DISCCONNECTION, DISCONNECTION_NAME, "")
	d.cl = true
	d.l = false
	d.c.Close()
}

func (d *Dial) On(name string, callback ClientEventCallback) {
	d.ad.On(name, callback)
}

func (d *Dial) Send(name string, data string) {
	d.send(_EVENT, name, data)
}