package handler

import (
	"flashdb/internal/storage"
	"fmt"
	"net"
	"strings"
)

const (
	PING = "PING"
	SET  = "SET"
	GET  = "GET"
	DEL  = "DEL"
)

type ConnectionHandler struct {
	storage *storage.Storage
}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{
		storage: storage.NewStorage(),
	}
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

		fmt.Println("Received message:", message)

		parts := strings.Fields(message)
		if len(parts) == 0 {
			continue
		}

		command := strings.ToUpper(parts[0])

		var response string
		switch command {
		case PING:
			response = "PONG\n"
		case SET:
			if len(parts) < 3 {
				response = "ERR wrong number of arguments for SET\n"
			} else {
				key := parts[1]
				value := strings.Join(parts[2:], " ")
				h.storage.Set(key, value)
				response = "OK\n"
			}
		case GET:
			if len(parts) < 2 {
				response = "ERR wrong number of arguments for GET\n"
			} else {
				key := parts[1]
				if value, exists := h.storage.Get(key); exists {
					response = value + "\n"
				} else {
					response = "(nil)\n"
				}
			}
		case DEL:
			if len(parts) < 2 {
				response = "ERR wrong number of arguments for DEL\n"
			} else {
				key := parts[1]
				if h.storage.Del(key) {
					response = "(integer) 1\n"
				} else {
					response = "(integer) 0\n"
				}
			}
		default:
			response = "UNKNOWN COMMAND\n"
		}

		conn.Write([]byte(response))
	}
}
