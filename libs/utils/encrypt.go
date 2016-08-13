package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
)

// Sha256 encrypts the given data using sha256
func Sha256(data string) string {
	h256 := sha256.New()
	io.WriteString(h256, data)
	return fmt.Sprintf("%x", h256.Sum(nil))
}
