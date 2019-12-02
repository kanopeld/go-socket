package core

import (
	"errors"
	"sync"
)

var (
	ErrRoomNotExist = errors.New("room not exist")
)

const (
	DefaultBroadcastRoomName = "defaultBroadcast"
)

type Broadcaster interface {
	Broadcast(event string, arg interface{}) error
}

type BroadcastAdaptor interface {
	Join(room string, c IdentifiableEmitter) error
	Leave(room string, c IdentifiableEmitter) error
	Send(ignore IdentifiableEmitter, room, event string, msg interface{}) error
	Len(room string) int
}

type rooms map[string]Room

type broadcast struct {
	rooms
	sync.RWMutex
}

func NewDefaultBroadcast() BroadcastAdaptor {
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
	var r, ok = b.rooms[room]
	b.RUnlock()
	if !ok {
		return ErrRoomNotExist
	}
	var err = r.RemoveClient(c)
	if err != nil {
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
