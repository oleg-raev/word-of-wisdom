package security

// ChallengeGenerator defines the interface for generating challenges.
type ChallengeGenerator interface {
	GenerateChallenge() (*Challenge, error)
}

// PoWValidator defines the interface for validating Proof of Work.
type PoWValidator interface {
	Validate(challenge, nonce string, complexity int) bool
}

// Challenge represents a PoW challenge including complexity.
type Challenge struct {
	Value      string
	Complexity int
}
