package security

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// DefaultChallengeGenerator implements ChallengeGenerator.
type DefaultChallengeGenerator struct {
	complexity int
}

// Ensure DefaultChallengeGenerator implements ChallengeGenerator.
var _ ChallengeGenerator = (*DefaultChallengeGenerator)(nil)

// NewChallengeGenerator creates a challenge generator with a fixed complexity.
func NewChallengeGenerator(complexity int) *DefaultChallengeGenerator {
	return &DefaultChallengeGenerator{
		complexity: complexity,
	}
}

// GenerateChallenge generates a new challenge with predefined complexity.
func (g *DefaultChallengeGenerator) GenerateChallenge() (*Challenge, error) {
	const challengeSize = 16 // Challenge size in bytes
	bytes := make([]byte, challengeSize)

	if _, err := rand.Read(bytes); err != nil {
		return nil, fmt.Errorf("failed to generate challenge: %w", err)
	}

	return &Challenge{
		Value:      hex.EncodeToString(bytes),
		Complexity: g.complexity,
	}, nil
}
