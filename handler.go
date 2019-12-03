package socket

import "sync"

type Events map[string]caller
type CallerMaker func(f interface{}) (caller, error)

type BaseHandler struct {
	Events
	*sync.RWMutex
	BroadcastAdaptor
	CallerMaker
}

func NewHandler(adaptor BroadcastAdaptor, maker CallerMaker) *BaseHandler {
	return &BaseHandler{
		Events:           make(Events),
		BroadcastAdaptor: adaptor,
		CallerMaker:      maker,
		RWMutex:          &sync.RWMutex{},
	}
}

func (h *BaseHandler) On(event string, f interface{}) error {
	c, err := h.CallerMaker(f)
	if err != nil {
		return err
	}
	h.Lock()
	h.Events[event] = c
	h.Unlock()
	return nil
}

func (h *BaseHandler) Off(event string) bool {
	h.Lock()
	_, ok := h.Events[event]
	delete(h.Events, event)
	h.Unlock()
	return ok
}

func (h *BaseHandler) GetEvents() Events {
	h.RLock()
	defer h.RUnlock()
	return h.Events
}

func (h *BaseHandler) GetBroadcast() BroadcastAdaptor {
	return h.BroadcastAdaptor
}
