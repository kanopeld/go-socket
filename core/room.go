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
	SetClient(c IdentifiableEmitter) error
	RemoveClient(client IdentifiableEmitter) error
	Len() int
	Send(ignore IdentifiableEmitter, event string, msg interface{}) error
	ClientExist(c IdentifiableEmitter) bool
}

type clients map[string]IdentifiableEmitter

type room struct {
	len int
	clients
	sync.RWMutex
}

func getRoom() Room {
	return &room{
		clients: make(clients, 0),
	}
}

func (r *room) SetClient(c IdentifiableEmitter) error {
	if r.ClientExist(c) {
		return ErrClientInRoomAlreadyExist
	}
	r.Lock()
	r.clients[c.ID()] = c
	r.len++
	r.Unlock()
	return nil
}

func (r *room) RemoveClient(client IdentifiableEmitter) error {
	r.RLock()
	_, ok := r.clients[client.ID()]
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

func (r *room) Send(ignore IdentifiableEmitter, event string, msg interface{}) error {
	r.Lock()
main:
	for _, c := range r.clients {
		if ignore != nil && ignore.ID() == c.ID() {
			continue main
		}
		_ = c.Emit(event, msg)
	}
	r.Unlock()
	return nil
}

func (r *room) ClientExist(c IdentifiableEmitter) (ok bool) {
	r.RLock()
	_, ok = r.clients[c.ID()]
	r.RUnlock()
	return ok
}
