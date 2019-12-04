package socket

import (
	"errors"
	"sync"
)

var (
	//ErrRoomNotExist will return if called room not exist
	ErrRoomNotExist = errors.New("room not exist")
)

const (
	//DefaultBroadcastRoomName name of default room in broadcast cluster. All new connections will be stored in room by this name
	DefaultBroadcastRoomName = "defaultBroadcast"
)

//Broadcaster organizes work with broadcast messaging
type Broadcaster interface {
	//Broadcast sends an event to the other side to everyone in the specified room
	Broadcast(event string, arg interface{}) error
}

//BroadcastAdaptor Available only on the server side.
//Organizes work with user associations in groups called "rooms".
//Serves to structure and identify possible zones of connected clients.
type BroadcastAdaptor interface {
	//Join adds the transferred client to the specified room. If the room does not exist, it will be created.
	Join(room string, c IdentifiableEmitter) error
	//Leave deletes the transferred client from the specified room. If after removal there are no clients left in the room, it will also be deleted
	Leave(room string, c IdentifiableEmitter) error
	//Send sends a message to all participants in the specified room
	Send(ignore IdentifiableEmitter, room, event string, msg interface{}) error
	//Len return rooms counter
	Len(room string) int
}

type rooms map[string]Room

type broadcast struct {
	rooms
	sync.RWMutex
}

func newDefaultBroadcast() BroadcastAdaptor {
	b := &broadcast{
		rooms: make(rooms, 0),
	}
	b.rooms[DefaultBroadcastRoomName] = getRoom()
	return b
}

func (b *broadcast) Join(room string, c IdentifiableEmitter) error {
	b.RLock()
	r, ok := b.rooms[room]
	b.RUnlock()
	if !ok {
		r = getRoom()
		b.Lock()
		b.rooms[room] = r
		b.Unlock()
	}
	if r.ClientExist(c) {
		return ErrClientInRoomAlreadyExist
	}
	return r.SetClient(c)
}

func (b *broadcast) Leave(room string, c IdentifiableEmitter) error {
	b.RLock()
	r, ok := b.rooms[room]
	b.RUnlock()
	if !ok {
		return ErrRoomNotExist
	}
	if err := r.RemoveClient(c); err != nil {
		return err
	}
	if r.Len() <= 0 {
		b.Lock()
		delete(b.rooms, room)
		b.Unlock()
	}
	return nil
}

func (b *broadcast) Send(ignore IdentifiableEmitter, room, event string, msg interface{}) error {
	b.Lock()
	r, ok := b.rooms[room]
	b.Unlock()
	if !ok {
		return ErrRoomNotExist
	}
	return r.Send(ignore, event, msg)
}

func (b *broadcast) Len(room string) int {
	b.Lock()
	r, ok := b.rooms[room]
	b.Unlock()
	if !ok {
		return -1
	}
	return r.Len()
}
