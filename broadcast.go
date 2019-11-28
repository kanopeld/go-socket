package socket

import "sync"

const (
	DefaultBroadcastRoomName = "defaultBroadcast"
)

type BroadcastAdaptor interface {
	Join(room string, c Client) error

	Leave(room string, c Client) error

	Send(ignore Client, room, event string, msg []byte) error

	Len(room string) int
}

type broadcast struct {
	rooms map[string]map[string]Client
	rMu   sync.Mutex
}

func newDefaultBroadcast() BroadcastAdaptor {
	b := &broadcast{
		rooms: make(map[string]map[string]Client),
	}

	b.rooms[DefaultBroadcastRoomName] = make(map[string]Client)

	return b
}

func (b *broadcast) Join(room string, c Client) error {
	b.rMu.Lock()
	r, ok := b.rooms[room]
	if !ok {
		r = make(map[string]Client)
		b.rooms[room] = r
	}

	r[c.ID()] = c
	b.rMu.Unlock()
	return nil
}

func (b *broadcast) Leave(room string, c Client) error {
	b.rMu.Lock()
	defer b.rMu.Unlock()

	r, ok := b.rooms[room]
	if !ok {
		return nil
	}

	delete(r, c.ID())
	if len(r) <= 0 {
		delete(b.rooms, room)
	}

	b.rooms[room] = r
	return nil
}

func (b *broadcast) Send(ignore Client, room, event string, msg []byte) error {
	b.rMu.Lock()
	defer b.rMu.Unlock()

	r, ok := b.rooms[room]
	if !ok {
		return nil
	}

	for id, c := range r {
		if id == ignore.ID() {
			continue
		}
		_ = c.Emit(event, msg)
	}

	return nil
}

func (b *broadcast) Len(room string) int {
	b.rMu.Lock()
	defer b.rMu.Unlock()

	r, ok := b.rooms[room]
	if !ok {
		return -1
	}

	return len(r)
}
