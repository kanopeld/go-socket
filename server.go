package socket

import (
	"net"
	"github.com/labstack/gommon/log"
	"sync"
)

type Server struct {
	*baseHandler
	ln net.Listener
	r bool
	sync.Mutex
}
func NewServer(p string) (*Server, error) {
	ln, err := net.Listen("tcp", p)
	if err != nil {
		return nil, err
	}

	s := &Server{
		baseHandler: newBaseHandler(newDefaultBroadcast()),
		ln:ln,
	}

	go s.loop()
	return s, nil
}

func (s *Server) loop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			continue
		}

		_, err = newClient(conn, s.baseHandler)
		if err != nil {
			log.Printf("Error create new client instance. Error: %s", err.Error())
		}
	}
}

func (s *Server) Start() {
	go s.loop()
	if s.r {
		return
	}

	s.r = true
	for s.r {

	}
}

func (s *Server) Stop() {
	s.r = false
	s.ln.Close()
}