package encodingtest

import (
	"github.com/mailchain/go-encoding"
)

// MustDecodeBase32 decodes a Base32 string
// It panics for invalid input.
func MustDecodeBase32(input string) []byte {
	dec, err := encoding.DecodeBase32(input)
	if err != nil {
		panic(err)
	}

	return dec
}
