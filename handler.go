package socket

import "sync"

type events map[string]caller
type CallerMaker func(f interface{}) (caller, error)

type BaseHandler struct {
	events
	sync.RWMutex
	BroadcastAdaptor
	CallerMaker
}

func NewHandler(adaptor BroadcastAdaptor, maker CallerMaker) *BaseHandler {
	return &BaseHandler{
		events:           make(events),
		BroadcastAdaptor: adaptor,
		CallerMaker:      maker,
		RWMutex:          sync.RWMutex{},
	}
}

func (h *BaseHandler) On(event string, f interface{}) error {
	c, err := h.CallerMaker(f)
	if err != nil {
		return err
	}
	h.Lock()
	h.events[event] = c
	h.Unlock()
	return nil
}

func (h *BaseHandler) Off(event string) bool {
	h.Lock()
	_, ok := h.events[event]
	delete(h.events, event)
	h.Unlock()
	return ok
}

func (h *BaseHandler) GetEvents() events {
	h.RLock()
	defer h.RUnlock()
	return h.events
}

func (h *BaseHandler) GetBroadcast() BroadcastAdaptor {
	return h.BroadcastAdaptor
}
