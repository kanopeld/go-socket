package socket

import "sync"

const (
	BASE_HANDLER_DEFAULT_NAME = "default"
)

type baseHandler struct {
	events map[string]*caller
	name string
	evMu sync.Mutex
}

func newBaseHandler() *baseHandler {
	return &baseHandler{
		events:make(map[string]*caller, 0),
		name:BASE_HANDLER_DEFAULT_NAME,
		evMu:sync.Mutex{},
	}
}

func (h *baseHandler) On(event string, f interface{}) error {
	c, err := newCaller(f)
	if err != nil {
		return err
	}
	h.evMu.Lock()
	h.events[event] = c
	h.evMu.Unlock()
	return nil
}

func (h *baseHandler) Off(event string) bool {
	h.evMu.Lock()
	_, ok := h.events[event]
	delete(h.events, event)
	h.evMu.Unlock()
	return ok
}

type clientHandler struct {
	*baseHandler
	client Client
}

func (h *clientHandler) call(event string, data []byte) error {
	h.evMu.Lock()
	c, ok := h.events[event]
	h.evMu.Unlock()

	if !ok {
		return nil
	}

	retV := c.Call(h.client, data)
	if len(retV) == 0 {
		return nil
	}

	var err error
	if last, ok := retV[len(retV)-1].Interface().(error); ok {
		err = last
		return err
	}

	return nil
}

func newClientHandler(c Client, bh *baseHandler) *clientHandler {
	events := make(map[string]*caller)
	bh.evMu.Lock()
	for k, v := range bh.events {
		events[k] = v
	}
	bh.evMu.Unlock()
	return &clientHandler{
		baseHandler: &baseHandler{
			events:    events,
			evMu:      bh.evMu,
		},
		client: c,
	}
}
