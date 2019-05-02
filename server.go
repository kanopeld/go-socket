package socket

import (
	"net"
	"sync"
)

type Server struct {
	*baseHandler
	ln net.Listener
	r bool
	sync.Mutex
}

func NewServer(p string, autostart bool) (*Server, error) {
	ln, err := net.Listen("tcp", p)
	if err != nil {
		return nil, err
	}

	s := &Server{
		baseHandler: newBaseHandler(newDefaultBroadcast()),
		ln:ln,
	}

	if autostart {
		s.r = true
		go s.loop()
	}
	return s, nil
}

func (s *Server) loop() {
	for s.r {
		conn, err := s.ln.Accept()
		if err != nil {
			continue
		}

		c, err := newClient(conn, s.baseHandler)
		if err != nil || c == nil {
			continue
		}
		c.loop()
	}
}

func (s *Server) Start() {
	if s.r {
		return
	}
	s.r = true
	go s.loop()
}

func (s *Server) Stop() {
	s.r = false
	_ = s.ln.Close()
}