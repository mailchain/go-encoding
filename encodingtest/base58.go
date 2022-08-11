package encodingtest

import (
	"github.com/mailchain/mailchain/encoding"
)

// MustDecodeBase58 decodes a Base58 string
// It panics for invalid input.
func MustDecodeBase58(input string) []byte {
	dec, err := encoding.DecodeBase58(input)
	if err != nil {
		panic(err)
	}

	return dec
}
