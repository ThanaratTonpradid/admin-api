package helper

import (
	"crypto/rand"
	"math/big"
)

func RandInt64(min, max int64) int64 {
	bg := big.NewInt(max - min + 1)
	n, _ := rand.Int(rand.Reader, bg)
	return n.Int64() + min
}

func RandFloat64() float64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(1<<62))
	return (float64(nBig.Int64()) / float64(1<<62))
}
