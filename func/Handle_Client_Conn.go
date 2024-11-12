package netcat

import (
	"fmt"
	"net"
)

func (s *Server) handleClientDisconnection(conn net.Conn, clientName string) {
	s.mu.Lock()
	delete(s.clients, conn)
	s.mu.Unlock()
	s.broadcastMessage(fmt.Sprintf("%s has left the chat", clientName), conn)
	fmt.Printf("Client Left The Server  ClientName: [%s] \n", clientName)
	s.messages = append(s.messages, fmt.Sprintf("%s has left the chat", clientName))
}

func (s *Server) broadcastMessage(msg string, sender net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for client := range s.clients {
		if client != sender {
			fmt.Fprintln(client, "\n"+msg)
			fmt.Fprint(client, HandleMessage("", s.clients[client]))
		}
	}
}
