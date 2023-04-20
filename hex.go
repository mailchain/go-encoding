package encoding

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// EncodeHexZeroX encodes src into "0x"+hex.Encode. As a convenience, it returns the encoding type used,
// but this value is always TypeHex0XPrefix.
// EncodeHexZeroX uses hexadecimal encoding prefixed with "0x".
func EncodeHexZeroX(src []byte) (encoded string) {
	out := make([]byte, len(src)*2+2)
	copy(out, "0x")
	hex.Encode(out[2:], src)

	return string(out)
}

// DecodeHexZeroX returns the bytes represented by the hexadecimal string src.
//
// DecodeHexZeroX expects that src contains only hex characters and must contain a `0x` prefix.
// If the input is malformed, DecodeHexZeroX returns an error.
func DecodeHexZeroX(in string) ([]byte, error) {
	if in == "" {
		return nil, fmt.Errorf("empty hex string")
	}

	if !strings.HasPrefix(in, "0x") {
		return nil, fmt.Errorf("missing \"0x\" prefix from hex string")
	}

	return hex.DecodeString(in[2:])
}

// EncodeHex returns the hexadecimal encoding of src.
func EncodeHex(src []byte) string {
	return hex.EncodeToString(src)
}

// DecodeHex returns the bytes represented by the hexadecimal string s.
//
// DecodeHex  expects that src contains only hexadecimal
// characters and that src has even length.
// If the input is malformed, DecodeString returns
// the bytes decoded before the error.
func DecodeHex(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
