package server

import (
	"flashdb/internal/handler"
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
	handler  *handler.ConnectionHandler
}

func NewServer(port string) (*Server, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
		handler:  handler.NewConnectionHandler(),
	}, nil
}

func (s *Server) Start() error {
	fmt.Printf("Server started on %s\n", s.listener.Addr())

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go s.handler.HandleConnection(conn)
	}
}

func (s *Server) Stop() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}
