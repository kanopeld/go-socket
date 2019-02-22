package socket

import (
	"net"
	"bufio"
	"strconv"
	"time"
	"math/rand"
)

type Client struct {
	conn net.Conn
	id string
	l bool
	c bool
	ad ClientAdaptor
	defAd Adaptor
	cC chan <- closeClient
}

func newClient(conn net.Conn, dad Adaptor, cC chan closeClient) (*Client, error) {
	nc := new(Client)
	nc.conn = conn
	nc.id = GetHash(strconv.Itoa(int(time.Now().Unix())) + strconv.FormatUint(rand.Uint64(), 10))
	nc.ad = GetClientAdapptor()
	nc.defAd = dad
	nc.cC = cC

	go nc.loop()

	return nc, nil
}

func (c *Client) loop() {
	defer func() {
		c.Close()
	}()

	if c.l || c.c {
		return
	}

	c.l = true
	for c.l {
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
			ev := c.defAd.GetEvent(CONNECTION_NAME)
			if ev == nil {
				continue
			}
			ev.callback(*c, "")
		case _PACKET_TYPE_DISCONNECT:
			c.ad.Call(DISCONNECTION_NAME, "")
			c.Close()
		case _PACKET_TYPE_EVENT:
			msg ,err := DecodeMessage(p.Payload)
			if err != nil {
				continue
			}

			c.ad.Call(msg.EventName, msg.Data.String())
		default:
			c.defAd.Call(ERROR_EVENT, "", c)
			c.defAd.Call(DISCONNECTION_NAME, "", c)
			c.Close()
		}
	}
}

func (c *Client) Close() {
	c.sendDisconnect()
	c.l = false
	c.c = true
	c.cC <- closeClient{id:c.id}
	c.conn.Close()
}

func (c *Client) ID() string {
	return c.id
}

func (c *Client) On(name string, callback ClientEventCallback) {
	c.ad.On(name, callback)
}

func (c *Client) send(t PackageType, name string, data string) {
	msg := Message{
		EventName:name,
		Data:MessagePayload{data:[]byte(data)},
	}

	p, err := NewEventPacket(msg)
	if err != nil {
		return
	}

	c.conn.Write(p.MarshalBinary())
}

func (c *Client) Send(name string, data string) {
	c.send(_PACKET_TYPE_EVENT, name, data)
}


func (c *Client) sendConnect() {
	p := NewPacket(_PACKET_TYPE_CONNECT)
	c.conn.Write(p.MarshalBinary())
}

func (c *Client) sendDisconnect() {
	p := NewPacket(_PACKET_TYPE_DISCONNECT)
	c.conn.Write(p.MarshalBinary())
}

func (c *Client) sendError() {
	p := NewPacket(_PACKET_TYPE_ERROR)
	c.conn.Write(p.MarshalBinary())
}