package socket

import "net"

type Sender interface {
	Send(p *Package) error
}

func getSender(c net.Conn) Sender {
	return &sender{c: c}
}

type sender struct {
	c net.Conn
}

func (s *sender) Send(p *Package) error {
	_, err := s.c.Write(p.MarshalBinary())
	return err
}
