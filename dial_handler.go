package socket

type dialHandler struct {
	*BaseHandler
	client DClient
}

func (h *dialHandler) call(event string, data []byte) error {
	h.hMu.RLock()
	c, ok := h.events[event]
	h.hMu.RUnlock()
	if !ok {
		return nil
	}
	retV := c.call(h.client, data)
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

func newDialHandler(c DClient) *dialHandler {
	ch := dialHandler{
		BaseHandler: NewHandler(nil, getCaller("DClient")),
		client:      c,
	}
	return &ch
}
