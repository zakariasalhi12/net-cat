package netcat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func (s *Server) handleNewClient(conn net.Conn) {
	defer conn.Close()
	fmt.Fprint(conn, string(LinuxLogo))
	fmt.Fprint(conn, "[ENTER YOUR NAME]: ")

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		clientName := strings.TrimSpace(scanner.Text())
		if clientName == "" {
			fmt.Fprintln(conn, "Name cannot be empty. Connection closed.")
			return
		}

		s.mu.Lock()
		if len(s.clients) >= 10 {
			s.mu.Unlock()
			fmt.Print("Chat room is full no one can Connect \n")
			fmt.Fprintln(conn, "Chat room is full. Connection closed.")
			return
		}
		s.clients[conn] = clientName
		s.mu.Unlock()
		fmt.Printf("New Client Join The Server ClientName: [%s] \n", clientName)
		s.broadcastMessage(fmt.Sprintf("%s has joined the chat", clientName), conn)
		s.messages = append(s.messages, fmt.Sprintf("%s has joined the chat", clientName))
		// Send all previous messages to the new client
		s.mu.Lock()
		for _, msg := range s.messages {
			fmt.Fprintln(conn, msg)
		}
		s.mu.Unlock()

		// Start listening for messages from this client
		s.listenForClientMessages(conn, clientName)
	}
}
