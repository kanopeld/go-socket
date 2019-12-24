package socket

import "sync"

type events map[string]HandlerCallback

type baseHandler struct {
	events
	hMu sync.RWMutex
	broadcastAdapter
}

func newHandler(adaptor broadcastAdapter) *baseHandler {
	return &baseHandler{
		events:           make(events),
		broadcastAdapter: adaptor,
		hMu:              sync.RWMutex{},
	}
}

// On registers an event handler under the given name.
func (h *baseHandler) On(event string, c HandlerCallback) {
	h.hMu.Lock()
	h.events[event] = c
	h.hMu.Unlock()
}

// Off deletes an event handler. Return true if event existed
func (h *baseHandler) Off(event string) bool {
	h.hMu.Lock()
	_, ok := h.events[event]
	delete(h.events, event)
	h.hMu.Unlock()
	return ok
}
