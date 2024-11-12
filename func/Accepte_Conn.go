package netcat

import "fmt"

func (s *Server) AcceptConnection() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go s.handleNewClient(conn)
	}
}
