package socket

import "sync"

type events map[string]caller
type callerMaker func(f interface{}) (caller, error)

type BaseHandler struct {
	events
	hMu sync.RWMutex
	BroadcastAdaptor
	callerMaker
}

func NewHandler(adaptor BroadcastAdaptor, maker callerMaker) *BaseHandler {
	return &BaseHandler{
		events:           make(events),
		BroadcastAdaptor: adaptor,
		callerMaker:      maker,
		hMu:              sync.RWMutex{},
	}
}

func (h *BaseHandler) On(event string, f interface{}) error {
	c, err := h.callerMaker(f)
	if err != nil {
		return err
	}
	h.hMu.Lock()
	h.events[event] = c
	h.hMu.Unlock()
	return nil
}

func (h *BaseHandler) Off(event string) bool {
	h.hMu.Lock()
	_, ok := h.events[event]
	delete(h.events, event)
	h.hMu.Unlock()
	return ok
}

func (h *BaseHandler) GetEvents() events {
	h.hMu.RLock()
	defer h.hMu.RUnlock()
	return h.events
}

func (h *BaseHandler) GetBroadcast() BroadcastAdaptor {
	return h.BroadcastAdaptor
}
