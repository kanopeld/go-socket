package socket

import (
	"net"
	"bufio"
	"strconv"
	"time"
)

type Client interface {
	ID() string

	Connection() net.Conn

	On(event string, f interface{}) error

	Off(event string) bool

	Emit(event string, data []byte) error

	Disconnect()
}

type client struct {
	*clientHandler
	conn net.Conn
	id string
}

func newClient(conn net.Conn, base *baseHandler) (Client, error) {
	nc := &client{
		conn: conn,
		id:GetHash(strconv.Itoa(int(time.Now().Unix())) + conn.RemoteAddr().String()),
	}
	nc.clientHandler = newClientHandler(nc, base)

	go nc.loop()
	return nc, nil
}

func (c *client) loop() {
	if err := c.send(& Package{PT:_PACKET_TYPE_CONNECT, Payload:[]byte(c.id)}); err != nil {
		c.Disconnect()
		return
	}

	for {
		msg, err := bufio.NewReader(c.conn).ReadBytes('\n')
		if err != nil {
			continue
		}

		p, err := DecodePackage(msg)
		if err != nil {
			continue
		}

		switch p.PT {
		case _PACKET_TYPE_CONNECT:
			if err := c.call(CONNECTION_NAME, nil); err != nil {
				c.Disconnect()
				return
			}
		case _PACKET_TYPE_DISCONNECT:
			c.Disconnect()
			return
		case _PACKET_TYPE_EVENT:
			msg ,err := DecodeMessage(p.Payload)
			if err != nil {
				continue
			}

			if err := c.call(msg.EventName, msg.Data); err != nil {
				c.Disconnect()
				return
			}
		}
	}
}

func (c *client) Emit(event string, data []byte) error {
	b, err := Message{EventName:event,Data:data}.MarshalBinary()
	if err != nil {
		return err
	}
	return c.send(&Package{PT:_PACKET_TYPE_EVENT, Payload:b})
}

func (c *client) ID() string {
	return c.id
}

func (c *client) Connection() net.Conn {
	return c.conn
}

func (c *client) Disconnect() {
	c.send(&Package{PT:_PACKET_TYPE_DISCONNECT})
	c.call(DISCONNECTION_NAME, nil)
	c.conn.Close()
}

func (c *client) send(p *Package) error {
	_, err := c.conn.Write(p.MarshalBinary())
	return err
}
