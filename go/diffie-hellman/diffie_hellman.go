package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

func PrivateKey(p *big.Int) *big.Int {
	var private big.Int

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	max := new(big.Int).Sub(p, big.NewInt(2))

	private.Rand(rnd, max)
	private.Add(&private, big.NewInt(2))

	return &private
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}