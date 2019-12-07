package socket

import (
	"errors"
	"sync"
)

var (
	// ErrClientNotInRoom informs the client is not in the room
	ErrClientNotInRoom = errors.New("client is not in this room")
	// ErrClientAlreadyInRoom informs the client is already in a room
	ErrClientAlreadyInRoom = errors.New("client is in the room already")
)

type clients map[string]IdentifiableEmitter

// Room serves to group of customers and work immediately with this group
type Room struct {
	len int
	clients
	sync.RWMutex
}

func getRoom() *Room {
	return &Room{
		clients: make(clients, 0),
	}
}

// SetClient adds a client to this room
func (r *Room) SetClient(c IdentifiableEmitter) error {
	if r.ClientExists(c) {
		return ErrClientAlreadyInRoom
	}
	r.Lock()
	r.clients[c.ID()] = c
	r.len++
	r.Unlock()
	return nil
}

// RemoveClient removes a client from this room
func (r *Room) RemoveClient(c IdentifiableEmitter) error {
	if !r.ClientExists(c) {
		return ErrClientNotInRoom
	}
	r.Lock()
	delete(r.clients, c.ID())
	r.len--
	r.Unlock()
	return nil
}

// Len returns amount of clients in this room
func (r *Room) Len() int {
	r.RLock()
	defer r.RUnlock()
	return r.len
}

// Send sends a message to the other all other clients
// It is possible to transfer the user to whom the message will not be transmitted
func (r *Room) Send(ignore IdentifiableEmitter, event string, msg []byte) error {
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

// ClientExists checks if a user is in this room
func (r *Room) ClientExists(c IdentifiableEmitter) (ok bool) {
	r.RLock()
	_, ok = r.clients[c.ID()]
	r.RUnlock()
	return ok
}
