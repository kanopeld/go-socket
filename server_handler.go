package socket

type serverHandler struct {
	*clientHandler
}

func (h *serverHandler) Broadcast(event string, msg interface{}) error {
	return h.BroadcastAdaptor.Send(h.client, DefaultBroadcastRoomName, event, msg)
}

func newServerHandler(c Client, bh handlerSharer) *serverHandler {
	return &serverHandler{
		clientHandler: &clientHandler{
			baseHandler: &baseHandler{
				events:           bh.getEvents(),
				BroadcastAdaptor: bh.getBroadcast(),
				callerMaker:      getCaller("SClient"),
			},
			client: c,
		},
	}
}
