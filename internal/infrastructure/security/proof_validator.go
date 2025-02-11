package security

import (
	"crypto/sha256"
	"math/big"
)

// SHA256PoWValidator implements PoWValidator using SHA-256.
type SHA256PoWValidator struct{}

// Ensure SHA256PoWValidator implements PoWValidator.
var _ PoWValidator = (*SHA256PoWValidator)(nil)

// Validate checks if the hash meets the required complexity.
func (s *SHA256PoWValidator) Validate(challenge, nonce string, complexity int) bool {
	hash := sha256.Sum256([]byte(challenge + nonce))
	hashInt := new(big.Int).SetBytes(hash[:])

	return hashInt.Rsh(hashInt, uint(256-complexity*4)).Uint64() == 0
}
