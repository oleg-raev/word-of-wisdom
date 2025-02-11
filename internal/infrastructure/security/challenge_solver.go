package security

import (
	"crypto/sha256"
	"math/big"
	"strconv"
)

type PoWSolver interface {
	Solve(challenge string, complexity int) string
}

// Ensure PoWSolverImpl implements PoWSolver.
var _ PoWSolver = (*PoWSolverImpl)(nil)

// PoWSolverImpl is responsible for solving Proof of Work challenges.
type PoWSolverImpl struct{}

// Solve finds a valid nonce for the given challenge and complexity.
func (s *PoWSolverImpl) Solve(challenge string, complexity int) string {
	target := new(big.Int).Lsh(big.NewInt(1), uint(256-complexity*4))
	hashInt := new(big.Int)
	for nonce := 0; ; nonce++ {
		hash := sha256.Sum256([]byte(challenge + strconv.Itoa(nonce)))
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			return strconv.Itoa(nonce)
		}
	}
}
