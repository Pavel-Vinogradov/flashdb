package handler

import (
	"fmt"
	"net"
	"strings"
)

const PING = "PING"

type ConnectionHandler struct{}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{}
}

func (h *ConnectionHandler) HandleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		message := string(buffer[:n])
		message = strings.TrimSpace(message)

		switch message {
		case PING:
			conn.Write([]byte("PONG\n"))
		default:
			conn.Write([]byte("UNKNOWN COMMAND\n"))
		}
		fmt.Println("Received message:", message)

		conn.Write([]byte("OK " + message))
	}
}
