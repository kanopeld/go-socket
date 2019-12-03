package socket

import (
	"sync"
)

type serverHandler struct {
	*BaseHandler
	client SClient
}

func (h *serverHandler) call(event string, data []byte) error {
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

func (h *serverHandler) Broadcast(event string, msg interface{}) error {
	return h.BroadcastAdaptor.Send(h.client, DefaultBroadcastRoomName, event, msg)
}

func newServerHandler(c SClient, bh HandlerSharer) *serverHandler {
	return &serverHandler{
		BaseHandler: &BaseHandler{
			Events:           bh.GetEvents(),
			BroadcastAdaptor: bh.GetBroadcast(),
			RWMutex:          &sync.RWMutex{},
			CallerMaker:      GetCaller("SClient"),
		},
		client: c,
	}
}
