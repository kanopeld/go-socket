package socket

import "sync"

type ClientStorage struct {
	list []*Client
	sync.Mutex
}

func NewClientStorage() *ClientStorage {
	ncs := new(ClientStorage)
	ncs.list = make([]*Client, 0)
	return ncs
}

func (cs *ClientStorage) Push(c *Client) (int) {
	cs.Lock()
	cs.list = append(cs.list, c)
	l := len(cs.list)
	cs.Unlock()

	return l
}

func (cs *ClientStorage) RemoveByID(id string) {
	cs.Lock()
	temp := make([]*Client, len(cs.list))
	for _, c := range cs.list {
		if c.id != id {
			temp = append(temp, c)
		}

		c.Close()
	}

	cs.list = temp
	cs.Unlock()
}