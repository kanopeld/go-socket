package socket

import (
	"errors"
	"sync"
)

var (
	// ErrRoomNotExist will return if called room not exist
	ErrRoomNotExist = errors.New("room not exist")
)

const (
	// DefaultBroadcastRoomName name of default room in broadcast cluster. All new connections will be stored in room by this name
	DefaultBroadcastRoomName = "defaultBroadcast"
)

type rooms map[string]*Room

// Broadcaster is available only on the server side.
// Organizes work with user associations in groups called "rooms".
// Serves to structure and identify possible zones of connected clients.
type Broadcaster struct {
	rooms
	sync.RWMutex
}

func newDefaultBroadcast() *Broadcaster {
	b := &Broadcaster{
		rooms: make(rooms, 0),
	}
	b.rooms[DefaultBroadcastRoomName] = getRoom()
	return b
}

// Join adds the transferred client to the specified room. If the room does not exist, it will be created.
func (b *Broadcaster) Join(room string, c IdentifiableEmitter) error {
	b.RLock()
	r, ok := b.rooms[room]
	b.RUnlock()
	if !ok {
		r = getRoom()
		b.Lock()
		b.rooms[room] = r
		b.Unlock()
	}
	if r.ClientExists(c) {
		return ErrClientAlreadyInRoom
	}
	return r.SetClient(c)
}

// Leave deletes the transferred client from the specified room. If after removal there are no clients left in the room, it will also be deleted
func (b *Broadcaster) Leave(room string, c IdentifiableEmitter) error {
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

// Send sends a message to all participants in the specified room
func (b *Broadcaster) Send(ignore IdentifiableEmitter, room, event string, msg []byte) error {
	b.Lock()
	r, ok := b.rooms[room]
	b.Unlock()
	if !ok {
		return ErrRoomNotExist
	}
	return r.Send(ignore, event, msg)
}

// Len return the amount of clients in a room
func (b *Broadcaster) Len(room string) int {
	b.Lock()
	r, ok := b.rooms[room]
	b.Unlock()
	if !ok {
		return -1
	}
	return r.Len()
}
