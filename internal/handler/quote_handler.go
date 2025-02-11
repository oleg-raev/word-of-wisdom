package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"word-of-wisdom/internal/infrastructure/security"
	"word-of-wisdom/internal/services/quotes"
)

// QuoteHandler contains the business logic for handling a connection.
type QuoteHandler struct {
	challengeGen security.ChallengeGenerator
	powValidator security.PoWValidator
	quoteService quotes.QuoteService
}

// NewQuoteHandler creates a new QuoteHandler.
func NewQuoteHandler(challengeGen security.ChallengeGenerator, powValidator security.PoWValidator, quoteService quotes.QuoteService) *QuoteHandler {
	return &QuoteHandler{
		challengeGen: challengeGen,
		powValidator: powValidator,
		quoteService: quoteService,
	}
}

// Handle processes a connection by performing the PoW challenge and sending a quote.
func (h *QuoteHandler) Handle(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Generate PoW challenge.
	challenge, err := h.challengeGen.GenerateChallenge()
	if err != nil {
		fmt.Fprintln(conn, "Error generating challenge")
		return
	}

	// Send challenge and complexity to the client.
	fmt.Fprintf(conn, "%s %d\n", challenge.Value, challenge.Complexity)

	// Read client's response (nonce).
	nonce, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(conn, "Error reading nonce")
		return
	}
	nonce = strings.TrimSpace(nonce)

	// Validate PoW solution.
	if !h.powValidator.Validate(challenge.Value, nonce, challenge.Complexity) {
		fmt.Fprintln(conn, "Invalid PoW")
		return
	}

	// Retrieve and send a random quote.
	quote, err := h.quoteService.GetRandomQuote()
	if err != nil {
		fmt.Fprintln(conn, "Error retrieving quotes")
		return
	}

	fmt.Fprintln(conn, quote)
}
