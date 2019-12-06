package socket

import "sync"

type events map[string]HandlerCallback

type baseHandler struct {
	events
	hMu sync.RWMutex
	BroadcastAdaptor
}

func newHandler(adaptor BroadcastAdaptor) *baseHandler {
	return &baseHandler{
		events:           make(events),
		BroadcastAdaptor: adaptor,
		hMu:              sync.RWMutex{},
	}
}

//On registers an event handler under the given name.
func (h *baseHandler) On(event string, c HandlerCallback) {
	h.hMu.Lock()
	h.events[event] = c
	h.hMu.Unlock()
}

//Off deletes an event handler. Return true if event was exist
func (h *baseHandler) Off(event string) bool {
	h.hMu.Lock()
	_, ok := h.events[event]
	delete(h.events, event)
	h.hMu.Unlock()
	return ok
}
