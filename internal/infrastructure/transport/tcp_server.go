package transport

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

// Handler defines a generic interface for handling a connection.
type Handler interface {
	Handle(conn net.Conn)
}

// TCPServer handles incoming TCP connections.
type TCPServer struct {
	address    string
	handler    Handler
	listener   net.Listener
	shutdownCh chan struct{}
}

// NewTCPServer creates a new TCPServer with the given address and handler.
func NewTCPServer(address string, handler Handler) *TCPServer {
	return &TCPServer{
		address:    address,
		handler:    handler,
		shutdownCh: make(chan struct{}),
	}
}

// Start begins listening for incoming connections.
func (s *TCPServer) Start() error {
	l, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	s.listener = l

	log.Printf("Server listening on %s", s.address)

	for {
		conn, err := l.Accept()
		if err != nil {
			// If shutdown is in progress, exit gracefully.
			select {
			case <-s.shutdownCh:
				log.Println("Shutdown signal received, stopping accept loop.")
				return nil
			default:
			}

			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Set connection timeout.
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		go s.handler.Handle(conn)
	}
}

// Shutdown gracefully shuts down the TCP server.
func (s *TCPServer) Shutdown(ctx context.Context) error {
	// Signal the accept loop to stop.
	close(s.shutdownCh)

	// Close the listener to unblock Accept.
	if s.listener != nil {
		return s.listener.Close()
	}

	return nil
}
