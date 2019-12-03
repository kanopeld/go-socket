package socket

type dialHandler struct {
	*BaseHandler
	client DClient
}

func (h *dialHandler) call(event string, data []byte) error {
	h.RLock()
	c, ok := h.Events[event]
	h.RUnlock()
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
		BaseHandler: NewHandler(nil, GetCaller("DClient")),
		client:      c,
	}
	return &ch
}
