package socket

import "sync"

type ClientStorage struct {
	list map[string]*Client
	sync.Mutex
}

func NewClientStorage() *ClientStorage {
	ncs := new(ClientStorage)
	ncs.list = make(map[string]*Client, 0)
	return ncs
}

func (cs *ClientStorage) Push(c *Client) (int) {
	cs.Lock()
	cs.list[c.ID()] = c
	l := len(cs.list)
	cs.Unlock()

	return l
}

func (cs *ClientStorage) Remove(id string) int {
	cs.Lock()
	delete(cs.list, id)
	l := len(cs.list)
	cs.Unlock()

	return l
}