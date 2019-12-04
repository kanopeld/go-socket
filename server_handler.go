package socket

type serverHandler struct {
	*baseHandler
	client SClient
}

func (h *serverHandler) call(event string, data []byte) error {
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

func (h *serverHandler) Broadcast(event string, msg interface{}) error {
	return h.BroadcastAdaptor.Send(h.client, DefaultBroadcastRoomName, event, msg)
}

func newServerHandler(c SClient, bh handlerSharer) *serverHandler {
	return &serverHandler{
		baseHandler: &baseHandler{
			events:           bh.getEvents(),
			BroadcastAdaptor: bh.getBroadcast(),
			callerMaker:      getCaller("SClient"),
		},
		client: c,
	}
}
