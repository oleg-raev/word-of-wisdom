package transport

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// TCPClient is responsible only for low-level TCP communication.
type TCPClient struct {
	serverAddress string // Address of the server to connect to
}

// NewTCPClient initializes a new TCP client.
func NewTCPClient(serverAddress string) *TCPClient {
	return &TCPClient{
		serverAddress: serverAddress,
	}
}

// Connect establishes a TCP connection to the server.
func (c *TCPClient) Connect() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.serverAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %v", err)
	}
	return conn, nil
}

// Send sends a message (terminated by newline) over the given connection.
func (c *TCPClient) Send(conn net.Conn, message string) error {
	fmt.Printf("Send %s\n", message)
	// Ensure the message is newline-terminated.
	trimmed := strings.TrimRight(message, "\n")
	_, err := fmt.Fprintln(conn, trimmed)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	return nil
}

// Receive reads a newline-terminated message from the connection.
func (c *TCPClient) Receive(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to receive message: %v", err)
	}
	fmt.Printf("Reveived %s\n", message)
	return message, nil
}
