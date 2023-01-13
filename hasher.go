package main

import (
	"crypto/sha256"
	"fmt"
)

type hash [32]byte

func newHash(str string) hash {
	return sha256.Sum256([]byte(str))
}

func (h *hash) String() string {
	return fmt.Sprintf("%x", h[:])
}
