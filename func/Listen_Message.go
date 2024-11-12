package netcat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func HandleMessage(Msg, ClientName string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s][%s]: %s", timestamp, ClientName, Msg)
}

func (s *Server) listenForClientMessages(conn net.Conn, clientName string) {
	fmt.Fprint(conn, HandleMessage("", clientName))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := strings.TrimSpace(scanner.Text())
		if msg == "" {
			fmt.Fprint(conn, HandleMessage("", clientName))
			continue
		}
		s.mu.Lock()
		fullMessage := HandleMessage(msg, clientName)
		s.messages = append(s.messages, fullMessage)
		fmt.Fprint(conn, HandleMessage("", clientName))
		s.mu.Unlock()
		s.broadcastMessage(fullMessage, conn)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from client:", err)
	}

	s.handleClientDisconnection(conn, clientName)
}
