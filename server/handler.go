package server

import (
	"github.com/kanopeld/go-socket/core"
)

type clientHandler struct {
	*core.BaseHandler
	client core.ServerClient
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

func (h *clientHandler) Broadcast(event string, msg interface{}) error {
	return h.BroadcastAdaptor.Send(h.client, core.DefaultBroadcastRoomName, event, msg)
}

func newClientHandler(c core.ServerClient, bh core.HandlerSharer) *clientHandler {
	return &clientHandler{
		BaseHandler: &core.BaseHandler{
			Events:           bh.GetEvents(),
			BroadcastAdaptor: bh.GetBroadcast(),
		},
		client: c,
	}
}
