package socket

import "net"

type Server struct {
	ln net.Listener
	cs *ClientStorage
	ad Adaptor
	clC bool
	closeClientChan chan closeClient
}

type closeClient struct {
	id string
}

func NewServer(p string) (s *Server, err error) {
	ln, err := net.Listen("tcp", p)
	if err != nil {
		return
	}

	s = &Server{
		ln:ln,
		cs:NewClientStorage(),
		ad:GetDefaultAdaptor(),
	}

	go s.loop()
	go s.observeClient()

	return
}

func (s *Server) loop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			continue
		}

		nc, err := newClient(conn, s.ad, s.closeClientChan)
		if err != nil {
			continue
		}

		s.cs.Push(nc)
	}
}

func (s *Server) observeClient() {
	for {
		select {
		case cC := <- s.closeClientChan:
			s.cs.Remove(cC.id)
		}
	}
}

func (s *Server) Close() {
	s.ln.Close()
}

func (s *Server) On(name string, callback EventCallback) {
	s.ad.On(name, callback)
}