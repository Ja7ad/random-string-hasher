package main

import (
	"crypto/sha256"
	"fmt"
)

func hashString(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
