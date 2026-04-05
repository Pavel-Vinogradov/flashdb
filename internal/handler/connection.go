package handler

import (
	"fmt"
	"net"
)

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
		fmt.Println("Received message:", message)

		conn.Write([]byte("OK" + message))
	}
}
