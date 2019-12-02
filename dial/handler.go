package dial

import (
	"github.com/kanopeld/go-socket/core"
)

type clientHandler struct {
	*core.BaseHandler
	client core.DClient
}

func (h *clientHandler) call(event string, data []byte) error {
	h.RLock()
	c, ok := h.Events[event]
	h.RUnlock()
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

func newClientHandler(c core.DClient) *clientHandler {
	ch := clientHandler{
		BaseHandler: core.NewHandler(nil, NewDialCaller),
		client:      c,
	}
	return &ch
}
