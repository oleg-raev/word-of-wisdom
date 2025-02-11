package quotes

import (
	"errors"
	"math/rand"
)

// Collection holds a set of quotes.
type Collection struct {
	quotes []string
}

// NewQuotesCollection creates a new Collection.
func NewQuotesCollection() *Collection {
	return &Collection{
		quotes: []string{},
	}
}

// AddQuote adds a quote to the collection.
func (qc *Collection) AddQuote(quote string) {
	qc.quotes = append(qc.quotes, quote)
}

// GetRandomQuote returns a random quote from the collection.
func (qc *Collection) GetRandomQuote() (string, error) {
	if len(qc.quotes) == 0 {
		return "", errors.New("no quotes available")
	}
	return qc.quotes[rand.Intn(len(qc.quotes))], nil
}
