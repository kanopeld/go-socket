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
	Len() uint
	Send(ignore []Client, event string, msg interface{}) error
}

type clients map[string]Client

type room struct {
	len uint
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
	var c, ok = r.clients[client.ID()]
	r.RUnlock()
	if !ok {
		ErrClientInRoomNotExist = errors.New(fmt.Sprintf("client with id (%s) in room not exist", c.ID()))
		return ErrClientInRoomNotExist
	}
	r.Lock()
	delete(r.clients, client.ID())
	r.len--
	r.Unlock()
	return nil
}

func (r *room) Len() uint {
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
