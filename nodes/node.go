package pg_query

import (
	"crypto/sha1"
	"hash"
)

type FingerprintContext struct {
	hash hash.Hash
}

func NewFingerprintContext() *FingerprintContext {
	return &FingerprintContext{hash: sha1.New()}
}

func (ctx FingerprintContext) Sum() []byte {
	return ctx.hash.Sum(nil)
}

type Node interface {
	Deparse() string
	Fingerprint(*FingerprintContext)
}
