package quotes

import (
	"bufio"
	"fmt"
	"os"

	"word-of-wisdom/internal/domain/quotes"
)

// QuoteService defines the operations for managing quotes.
type QuoteService interface {
	LoadFromFile(filename string) error
	GetRandomQuote() (string, error)
}

// QuoteServiceImpl implements QuoteService with an in-memory quote collection.
type QuoteServiceImpl struct {
	collection *quotes.Collection
}

// Ensure QuoteServiceImpl implements QuoteService.
var _ QuoteService = (*QuoteServiceImpl)(nil)

// NewQuoteService creates a new QuoteServiceImpl.
func NewQuoteService() *QuoteServiceImpl {
	return &QuoteServiceImpl{
		collection: quotes.NewQuotesCollection(),
	}
}

// LoadFromFile reads quotes from a file and stores them in the collection.
func (s *QuoteServiceImpl) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s.collection.AddQuote(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

// GetRandomQuote retrieves a random quote from the collection.
func (s *QuoteServiceImpl) GetRandomQuote() (string, error) {
	return s.collection.GetRandomQuote()
}
