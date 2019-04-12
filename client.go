package socket

import (
	"net"
	"bufio"
	"strconv"
	"time"
	"crypto/md5"
	"encoding/hex"
)

type Client interface {
	ID() string

	Connection() net.Conn

	On(event string, f interface{}) error

	Off(event string) bool

	Emit(event string, data []byte) error

	Broadcast(event string, msg []byte) error

	Disconnect()
}

type client struct {
	*clientHandler
	*defaultEmitter
	conn net.Conn
	id string
}

func newClient(conn net.Conn, base *baseHandler) (Client, error) {
	nc := &client{
		conn: conn,
		id:newID(conn),
		defaultEmitter:&defaultEmitter{c: conn},
	}
	nc.clientHandler = newClientHandler(nc, base)
	err := nc.baseHandler.broadcast.Join(DefaultBroadcastRoomName, nc)
	if err != nil {
		return nil, err
	}

	go nc.loop()
	return nc, nil
}

func (c *client) loop() {
	defer func() {
		c.Disconnect()
		_ = c.broadcast.Leave(DefaultBroadcastRoomName, c)
	}()

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
			if err := c.send(& Package{PT:_PACKET_TYPE_CONNECT, Payload:[]byte(c.id)}); err != nil {
				c.Disconnect()
				return
			}

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
	_ = c.send(&Package{PT:_PACKET_TYPE_DISCONNECT})
	_ = c.call(DISCONNECTION_NAME, nil)
	_ = c.conn.Close()
}

func newID(c net.Conn) string {
	st := strconv.Itoa(int(time.Now().Unix())) + c.RemoteAddr().String()
	hasher := md5.New()
	hasher.Write([]byte(st))
	hash := hex.EncodeToString(hasher.Sum(nil)[:16])
	return hash
}
