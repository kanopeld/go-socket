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

		dec, err := NewMessageDecoder(msg)
		if err != nil {
			continue
		}

		switch dec.mt {
		case _CONNECTION:
			ev := c.defAd.GetEvent(CONNECTION_NAME)
			if ev == nil {
				continue
			}
			ev.callback(*c, "")
		case _EVENT:
			c.ad.Call(dec.eventName, dec.payload)
		case _ERROR:
			c.defAd.Call(DISCONNECTION_NAME, "", c)
		case _DISCCONNECTION:
			c.ad.Call(DISCONNECTION_NAME, "")
			return
		}
	}
}

func (c *Client) Close() {
	c.send(_DISCCONNECTION, DISCONNECTION_NAME, "")
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

func (c *Client) send(t int, name string, data string) error {
	msg, err := NewEncodeMessage(t, name, data)
	if err != nil {
		return err
	}

	_, err = c.conn.Write([]byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Send(name string, data string) {
	c.send(_EVENT, name, data)
}
