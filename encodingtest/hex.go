package encodingtest

import (
	"github.com/mailchain/encoding"
)

// MustDecodeHex decodes a hex string. It panics for invalid input.
func MustDecodeHex(input string) []byte {
	dec, err := encoding.DecodeHex(input)
	if err != nil {
		panic(err)
	}

	return dec
}
