package encodingtest

import "github.com/mailchain/encoding"

// MustDecodeHexZeroX decodes a hex string. It panics for invalid input.
func MustDecodeHexZeroX(in string) []byte {
	dec, err := encoding.DecodeHexZeroX(in)
	if err != nil {
		panic(err)
	}

	return dec
}
