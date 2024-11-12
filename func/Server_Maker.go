package netcat

import (
	"net"
	"sync"
)

var LinuxLogo []byte

type Server struct {
	serverAddr string
	ln         net.Listener
	clients    map[net.Conn]string
	mu         sync.Mutex
	messages   []string
}

func MakeServer(port string) *Server {
	return &Server{
		serverAddr: ":" + port,
		clients:    make(map[net.Conn]string),
		messages:   make([]string, 0),
	}
}
