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

func newClient(conn net.Conn, base core.HandlerSharer) (core.Looper, error) {
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
	_, _ = c.conn.Write(core.Package{PT: core.PackTypeConnect, Payload: []byte(c.id)}.MarshalBinary())
}

func (c *client) Loop() {
	defer func() {
		c.Disconnect()
		_ = c.Leave(core.DefaultBroadcastRoomName, c)
	}()

	c.sendConnect()
	reader := bufio.NewReader(c.conn)
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
		case core.PackTypeConnectAccept:
			if err := c.call(core.ConnectionName, nil); err != nil {
				return
			}
		case core.PackTypeDisconnect:
			return
		case core.PackTypeEvent:
			msg := core.DecodeMessage(p.Payload)
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
	_ = c.Send(&core.Package{PT: core.PackTypeDisconnect})
	_ = c.call(core.DisconnectionName, nil)
	_ = c.conn.Close()
}

func newID(c net.Conn) string {
	return strconv.Itoa(int(time.Now().Unix())) + c.RemoteAddr().String()
}
