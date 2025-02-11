package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"word-of-wisdom/internal/infrastructure/security"
	"word-of-wisdom/internal/infrastructure/transport"
)

func main() {
	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		log.Fatalf("Error getting SERVER_ADDR parameter")
	}

	// Create a low-level TCP client from the transport layer.
	tcpClient := transport.NewTCPClient(serverAddr)

	// Connect to the server.
	conn, err := tcpClient.Connect()
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	defer conn.Close()

	// Step 1: Receive the PoW challenge and complexity from the server.
	challengeMsg, err := tcpClient.Receive(conn)
	if err != nil {
		log.Fatalf("Error receiving challenge: %v", err)
	}

	challenge, complexity, err := parseChallengeMessage(challengeMsg)
	if err != nil {
		log.Fatalf("Error parsing challenge: %v", err)
	}

	// Step 2: Solve the PoW challenge using the PoWSolver from the security package.
	solver := &security.PoWSolverImpl{}
	nonce := solver.Solve(challenge, complexity)

	// Step 3: Send the solved nonce back to the server.
	if err := tcpClient.Send(conn, nonce); err != nil {
		log.Fatalf("Error sending nonce: %v", err)
	}

	// Step 4: Receive the quote from the server.
	quote, err := tcpClient.Receive(conn)
	//todo handle errors properly (now we do not know if it's error or quote)
	if err != nil {
		log.Fatalf("Error receiving quote: %v", err)
	}

	fmt.Printf("Received Quote: %s\n", strings.TrimSpace(quote))
}

// parseChallengeMessage extracts the challenge and complexity from the server's message.
// Expected format: "<challenge> <complexity>"
func parseChallengeMessage(message string) (string, int, error) {
	message = strings.TrimSpace(message)
	idx := strings.LastIndexByte(message, ' ')
	if idx == -1 {
		return "", 0, fmt.Errorf("invalid challenge format: expected 2 parts separated by space")
	}

	challenge := message[:idx]
	complexity, err := strconv.Atoi(message[idx+1:])
	if err != nil {
		return "", 0, fmt.Errorf("invalid complexity value: %v", err)
	}

	return challenge, complexity, nil
}
