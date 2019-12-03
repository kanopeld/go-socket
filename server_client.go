package socket

import (
	"bufio"

	"net"
	"strconv"
	"time"
)

type client struct {
	*serverHandler
	Emitter
	conn net.Conn
	id   string
	disc bool
}

func newClient(conn net.Conn, base HandlerSharer) (Looper, error) {
	nc := &client{
		conn:    conn,
		id:      newID(conn),
		Emitter: GetEmitter(conn),
	}
	nc.serverHandler = newServerHandler(nc, base)
	err := nc.Join(DefaultBroadcastRoomName, nc)
	if err != nil {
		_ = nc.Leave(DefaultBroadcastRoomName, nc)
		return nil, err
	}
	return nc, nil
}

func (c *client) sendConnect() {
	_, _ = c.conn.Write(Package{PT: PackTypeConnect, Payload: []byte(c.id)}.MarshalBinary())
}

func (c *client) Loop() {
	defer func() {
		c.Disconnect()
		_ = c.Leave(DefaultBroadcastRoomName, c)
	}()

	c.sendConnect()
	reader := bufio.NewReader(c.conn)
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
		case PackTypeConnectAccept:
			if err := c.call(ConnectionName, nil); err != nil {
				return
			}
		case PackTypeDisconnect:
			return
		case PackTypeEvent:
			msg := decodeMessage(p.Payload)
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
	_ = c.Send(&Package{PT: PackTypeDisconnect})
	_ = c.call(DisconnectionName, nil)
	_ = c.conn.Close()
}

func newID(c net.Conn) string {
	return strconv.Itoa(int(time.Now().Unix())) + c.RemoteAddr().String()
}
