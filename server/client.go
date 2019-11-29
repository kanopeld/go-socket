package server

import (
	"bufio"
	"github.com/kanopeld/go-socket/core"
	"net"
	"strconv"
	"time"
)

type client struct {
	*clientHandler
	core.Emitter
	conn net.Conn
	id   string
	disc bool
}

func newClient(conn net.Conn, base core.HandlerSharer) (looper, error) {
	nc := &client{
		conn:    conn,
		id:      newID(conn),
		Emitter: core.GetEmitter(conn),
	}
	nc.clientHandler = newClientHandler(nc, base)
	err := nc.Join(core.DefaultBroadcastRoomName, nc)
	if err != nil {
		_ = nc.Leave(core.DefaultBroadcastRoomName, nc)
		return nil, err
	}
	return nc, nil
}

func (c *client) sendConnect() {
	_, _ = c.conn.Write(Package{PT: _PACKET_TYPE_CONNECT, Payload: []byte(c.id)}.MarshalBinary())
}

func (c *client) loop() {
	defer func() {
		c.Disconnect()
		_ = c.broadcast.Leave(core.DefaultBroadcastRoomName, c)
	}()

	c.sendConnect()
	reader := bufio.NewReader(c.conn)
	for {
		msg, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		p, err := DecodePackage(msg)
		if err != nil {
			return
		}

		switch p.PT {
		case _PACKET_TYPE_CONNECT_ACCEPT:
			if err := c.call(CONNECTION_NAME, nil); err != nil {
				return
			}
		case _PACKET_TYPE_DISCONNECT:
			return
		case _PACKET_TYPE_EVENT:
			msg := DecodeMessage(p.Payload)
			if err := c.call(msg.EventName, msg.Data); err != nil {
				return
			}
		}
	}
}

func (c *client) ID() string {
	return c.id
}

func (c *client) Connection() net.Conn {
	return c.conn
}

func (c *client) Disconnect() {
	if c.disc {
		return
	}
	c.disc = true
	_ = c.send(&Package{PT: _PACKET_TYPE_DISCONNECT})
	_ = c.call(DISCONNECTION_NAME, nil)
	_ = c.conn.Close()
}

func newID(c net.Conn) string {
	return strconv.Itoa(int(time.Now().Unix())) + c.RemoteAddr().String()
}
