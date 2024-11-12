package netcat

import "net"

func (s *Server) StartServer() error {
	ln, err := net.Listen("tcp4", s.serverAddr)
	if err != nil {
		return err
	}
	s.ln = ln
	go s.AcceptConnection() // Start accepting connections in a separate goroutine
	return nil
}
