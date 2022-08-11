package encoding

import "github.com/mr-tron/base58"

// DecodeBase58 returns the bytes represented by the base58 string src.
//
// DecodeBase58 expects that src contains only base58 characters.
// If the input is malformed, DecodeBase58 returns an error.
func DecodeBase58(src string) ([]byte, error) {
	return base58.Decode(src)
}

// EncodeBase58 returns the string represented by the base58 byte src.
//
// EncodeBase58 expects that src contains only base58 byte.
// If the input is malformed, EncodeBase58 returns an error.
func EncodeBase58(src []byte) string {
	return base58.Encode(src)
}
