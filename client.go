package socket

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"net"
	"strconv"
	"time"
)

// Client controls the client side of a connection
type Client interface {
	ID() string

	Connection() net.Conn

	On(event string, f interface{}) error

	Off(event string) bool

	Emit(event string, arg interface{}) error

	Broadcast(event string, msg []byte) error

	Disconnect()
}

type client struct {
	*clientHandler
	*defaultEmitter
	conn net.Conn
	id   string
	disc bool
}

func newClient(conn net.Conn, base *baseHandler) (looper, error) {
	nc := &client{
		conn:           conn,
		id:             newID(conn),
		defaultEmitter: &defaultEmitter{c: conn},
	}
	nc.clientHandler = newClientHandler(nc, base)
	err := nc.baseHandler.broadcast.Join(DefaultBroadcastRoomName, nc)
	if err != nil {
		_ = base.broadcast.Leave(DefaultBroadcastRoomName, nc)
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
		_ = c.broadcast.Leave(DefaultBroadcastRoomName, c)
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
	st := strconv.Itoa(int(time.Now().Unix())) + c.RemoteAddr().String()
	hasher := md5.New()
	hasher.Write([]byte(st))
	hash := hex.EncodeToString(hasher.Sum(nil)[:16])
	return hash
}
