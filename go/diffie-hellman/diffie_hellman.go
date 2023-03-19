package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

// PrivateKey generates and returns a new private key
// between 1 and p (extremes excluded)
func PrivateKey(p *big.Int) *big.Int {
	min := big.NewInt(2)
	max := big.NewInt(0).Sub(p, min)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil
	}
	return big.NewInt(0).Add(n, min)
}

// PublicKey generates a new public from the given private key, p and g values
// using the formula  A = g**private mod p
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

// NewPair computes and returns a private key and a public key
// for the given p and g values
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

// SecretKey returns a secret key given p, one's own private key
// and another's public key generated with the same p value (a prime number)
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
