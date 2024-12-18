package main

import (
	"fmt"
	"os"

	netcat "netcat/func"
)

var Port = "8989" // Default Port

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(os.Args) == 2 {
		Port = os.Args[1]
	}

	netcat.LinuxLogo, _ = os.ReadFile("LinuxLogo.txt")
	newServer := netcat.MakeServer(Port)
	err := newServer.StartServer()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Printf("Server Started At port :%s\n", Port)
	select {} // Keep the server running
}
