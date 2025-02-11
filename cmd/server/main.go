package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"word-of-wisdom/internal/config"
	"word-of-wisdom/internal/handler"
	"word-of-wisdom/internal/infrastructure/security"
	"word-of-wisdom/internal/infrastructure/transport"
	"word-of-wisdom/internal/services/quotes"
)

func main() {
	// Load configuration from environment variables.
	cfg := config.LoadConfig()

	// Initialize PoW services using config values.
	challengeGenerator := security.NewChallengeGenerator(cfg.ChallengeComplexity)
	powValidator := &security.SHA256PoWValidator{}

	// Initialize the QuoteService and load quotes from the configured file.
	quoteService := quotes.NewQuoteService()
	if err := quoteService.LoadFromFile(cfg.QuotesFile); err != nil {
		log.Fatalf("Failed to load quotes: %v", err)
	}

	// Create the QuoteHandler which contains the business logic.
	quoteHandler := handler.NewQuoteHandler(challengeGenerator, powValidator, quoteService)

	// Create the TCP server and pass the handler, using the server address from config.
	server := transport.NewTCPServer(cfg.ServerAddr, quoteHandler)

	// Set up channel to capture system signals for graceful shutdown.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the server in a separate goroutine.
	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	log.Println("Server started. Press Ctrl+C to stop.")

	// Wait for a termination signal.
	<-stop
	log.Println("Shutting down server...")

	// Shutdown the server gracefully.
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error during shutdown: %v", err)
	}

	log.Println("Server stopped gracefully.")
}
