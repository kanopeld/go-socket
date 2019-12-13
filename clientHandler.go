package socket

import "errors"

var (
	// ErrEventNotExist is returned if the emited event does not exist
	ErrEventNotExist = errors.New("event does not exist")
)

type clientHandler struct {
	*baseHandler
	client Client
}

func (h *clientHandler) call(event string, data []byte) error {
	h.hMu.RLock()
	c, ok := h.events[event]
	h.hMu.RUnlock()
	if !ok {
		return ErrEventNotExist
	}
	return c(h.client, data)
}

func (h *clientHandler) Broadcast(event string, msg []byte) error {
	if h.Broadcaster == nil {
		return nil
	}
	return h.Send(h.client, DefaultBroadcastRoomName, event, msg)
}

func newClientHandler(c Client, bh *baseHandler) *clientHandler {
	return &clientHandler{
		client:      c,
		baseHandler: bh,
	}
}
