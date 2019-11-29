package core

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrClientInRoomNotExist error
)

type Room interface {
	SetClient(c Client) Room
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

func (r *room) SetClient(c Client) Room {
	r.Lock()
	r.clients[c.ID()] = c
	r.len++
	r.Unlock()
	return r
}

func (r *room) RemoveClient(client Client) error {
	r.RLock()
	var _, ok = r.clients[client.ID()]
	r.RUnlock()
	if !ok {
		ErrClientInRoomNotExist = errors.New(fmt.Sprintf("client with id (%s) in room not exist", client.ID()))
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
