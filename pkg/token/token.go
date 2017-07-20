package token

import (
	"crypto/rand"
	"fmt"
)

func Rand() string {
	b := make([]byte, 8)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}
