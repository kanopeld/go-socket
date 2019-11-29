package core

import (
	"errors"
	"sync"
)

var (
	ErrClientInRoomNotExist     = errors.New("client in room not exist")
	ErrClientInRoomAlreadyExist = errors.New("client not exist in room")
)

type Room interface {
	SetClient(c Client) error
	RemoveClient(client Client) error
	Len() int
	Send(ignore []Client, event string, msg interface{}) error
	ClientExist(c Client) bool
}

type clients map[string]Client

type room struct {
	len int
	clients
	sync.RWMutex
}

func getRoom() *room {
	return &room{
		clients: make(clients, 0),
	}
}

func (r *room) SetClient(c Client) error {
	if r.ClientExist(c) {
		return ErrClientInRoomAlreadyExist
	}
	r.Lock()
	r.clients[c.ID()] = c
	r.len++
	r.Unlock()
	return nil
}

func (r *room) RemoveClient(client Client) error {
	r.RLock()
	var _, ok = r.clients[client.ID()]
	r.RUnlock()
	if !ok {
		return ErrClientInRoomNotExist
	}
	r.Lock()
	delete(r.clients, client.ID())
	r.len--
	r.Unlock()
	return nil
}

func (r *room) Len() int {
	r.RLock()
	defer r.RUnlock()
	return r.len
}

func (r *room) Send(ignore []Client, event string, msg interface{}) error {
	r.Lock()
main:
	for _, c := range r.clients {
		if ignore != nil {
			for _, ic := range ignore {
				if ic.ID() == c.ID() {
					continue main
				}
			}
		}
		_ = c.Emit(event, msg)
	}
	r.Unlock()
	return nil
}

func (r *room) ClientExist(c Client) bool {
	r.RLock()
	var _, ok = r.clients[c.ID()]
	r.RUnlock()
	return ok
}
