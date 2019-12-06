package socket

import (
	"errors"
	"sync"
)

var (
	//ErrClientInRoomNotExist error will return when you try to contact a user who is not in the specified room
	ErrClientInRoomNotExist = errors.New("client in room not exist")
	//ErrClientInRoomAlreadyExist error will return when you try to re-record the user in the room in which he already exists
	ErrClientInRoomAlreadyExist = errors.New("client in room alrea")
)

//Room serves to group customers and work immediately with this group.
type Room interface {
	//SetClient adds a client to this room.
	SetClient(c IdentifiableEmitter) error
	//RemoveClient removes a client from this room.
	RemoveClient(client IdentifiableEmitter) error
	//Len will return current client count from this room.
	Len() int
	//Send sends a message to the other side of all clients of this contact.
	//It is possible to transfer the user to whom the message will not be transmitted
	Send(ignore IdentifiableEmitter, event string, msg []byte) error
	//ClientExist check that given client exist in this room
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

//SetClient adds a client to this room.
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

//RemoveClient removes a client from this room.
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

//Len will return current client count from this room.
func (r *room) Len() int {
	r.RLock()
	defer r.RUnlock()
	return r.len
}

//Send sends a message to the other side of all clients of this contact.
//It is possible to transfer the user to whom the message will not be transmitted
func (r *room) Send(ignore IdentifiableEmitter, event string, msg []byte) error {
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

//ClientExist check that given client exist in this room
func (r *room) ClientExist(c IdentifiableEmitter) (ok bool) {
	r.RLock()
	_, ok = r.clients[c.ID()]
	r.RUnlock()
	return ok
}
