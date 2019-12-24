package socket

import (
	"sync"
)

const (
	// DefaultBroadcastRoomName name of default room in broadcast cluster. All new connections will be stored in room by this name
	DefaultBroadcastRoomName = "defaultBroadcast"
)

type rooms map[string]roomer

// Broadcast is available only on the server side.
// Organizes work with user associations in groups called "rooms".
// Serves to structure and identify possible zones of connected clients.
type broadcast struct {
	rooms
	sync.RWMutex
}

func newDefaultBroadcast() broadcastAdapter {
	b := &broadcast{
		rooms: make(rooms, 0),
	}
	b.rooms[DefaultBroadcastRoomName] = getRoom()
	return b
}

// Join adds the transferred client to the specified room. If the room does not exist, it will be created.
func (b *broadcast) join(room string, c Client) error {
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
func (b *broadcast) leave(room string, c Client) error {
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
func (b *broadcast) send(ignore Client, room, event string, msg []byte) error {
	b.Lock()
	r, ok := b.rooms[room]
	b.Unlock()
	if !ok {
		return ErrRoomNotExist
	}
	return r.Send(ignore, event, msg)
}

// Len return the amount of clients in a room
func (b *broadcast) len(room string) int {
	b.Lock()
	r, ok := b.rooms[room]
	b.Unlock()
	if !ok {
		return -1
	}
	return r.Len()
}
