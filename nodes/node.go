package pg_query

import (
	"crypto/sha1"
	"hash"
	"io"
)

type FingerprintContext interface {
	WriteString(string)
}

type FingerprintHashContext struct {
	hash hash.Hash
}

func NewFingerprintHashContext() *FingerprintHashContext {
	return &FingerprintHashContext{hash: sha1.New()}
}

func (ctx FingerprintHashContext) WriteString(str string) {
	io.WriteString(ctx.hash, str)
}

func (ctx FingerprintHashContext) Sum() []byte {
	return ctx.hash.Sum(nil)
}

type Node interface {
	Deparse() string
	Fingerprint(FingerprintContext)
}
