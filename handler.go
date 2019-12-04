package socket

import "sync"

type events map[string]caller
type callerMaker func(f interface{}) (caller, error)

type baseHandler struct {
	events
	hMu sync.RWMutex
	BroadcastAdaptor
	callerMaker
}

func newHandler(adaptor BroadcastAdaptor, maker callerMaker) *baseHandler {
	return &baseHandler{
		events:           make(events),
		BroadcastAdaptor: adaptor,
		callerMaker:      maker,
		hMu:              sync.RWMutex{},
	}
}

func (h *baseHandler) On(event string, f interface{}) error {
	c, err := h.callerMaker(f)
	if err != nil {
		return err
	}
	h.hMu.Lock()
	h.events[event] = c
	h.hMu.Unlock()
	return nil
}

func (h *baseHandler) Off(event string) bool {
	h.hMu.Lock()
	_, ok := h.events[event]
	delete(h.events, event)
	h.hMu.Unlock()
	return ok
}

func (h *baseHandler) getEvents() events {
	h.hMu.RLock()
	defer h.hMu.RUnlock()
	return h.events
}

func (h *baseHandler) getBroadcast() BroadcastAdaptor {
	return h.BroadcastAdaptor
}
