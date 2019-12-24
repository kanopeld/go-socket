package socket

import (
	"bufio"

	"net"
	"strconv"
	"time"
)

type client struct {
	*clientHandler
	Emitter
	conn net.Conn
	id   string
	disc bool
}

func newClient(conn net.Conn, base *baseHandler) (looper, error) {
	nc := &client{
		conn:    conn,
		Emitter: getEmitter(conn),
	}
	nc.setNewID()
	nc.clientHandler = newClientHandler(nc, base)
	err := nc.join(DefaultBroadcastRoomName, nc)
	if err != nil {
		_ = nc.leave(DefaultBroadcastRoomName, nc)
		return nil, err
	}
	return nc, nil
}

func (c *client) sendConnect() {
	_, _ = c.conn.Write(sockPackage{PT: PackTypeConnect, Payload: []byte(c.id)}.MarshalBinary())
}

func (c *client) loop() {
	defer func() {
		c.Disconnect()
		_ = c.leave(DefaultBroadcastRoomName, c)
	}()

	c.sendConnect()
	reader := bufio.NewReader(c.conn)
	for {
		msg, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		p := decodePackage(msg)
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
	_ = c.send(&sockPackage{PT: PackTypeDisconnect})
	_ = c.call(DisconnectionName, nil)
	_ = c.conn.Close()
}

func (c *client) setNewID() {
	c.id = strconv.Itoa(int(time.Now().Unix())) + c.conn.RemoteAddr().String()
}
