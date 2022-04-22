package encoding

import (
	"errors"
	"strings"
)

var errUnsupportedEncoding = errors.New("encoding not supported")

// Decode returns the bytes represented by the decoded string src.
//
// Decode uses the decode method mapped to kind parameter.
// If the input is kind is unknown or the input is malformed for the decode method it returns an error.
func Decode(encoding, src string) ([]byte, error) {
	switch strings.ToLower(encoding) {
	case KindBase64:
		return DecodeBase64(src)
	case KindBase64URL:
		return DecodeBase64URL(src)
	case KindBase58:
		return DecodeBase58(src)
	case KindBase32:
		return DecodeBase32(src)
	case KindHex:
		return DecodeHex(src)
	case KindHex0XPrefix:
		return DecodeHexZeroX(src)
	case KindMnemonicAlgorand:
		return DecodeMnemonicAlgorand(src)
	default:
		return nil, errUnsupportedEncoding
	}
}

// Encode returns the bytes encoded as requested by the encoding parameter.
func Encode(encoding string, src []byte) (string, error) {
	switch strings.ToLower(encoding) {
	case KindBase64URL:
		return EncodeBase64URL(src), nil
	case KindBase64:
		return EncodeBase64(src), nil
	case KindBase58:
		return EncodeBase58(src), nil
	case KindBase32:
		return EncodeBase32(src), nil
	case KindHex:
		return EncodeHex(src), nil
	case KindHex0XPrefix:
		return EncodeHexZeroX(src), nil
	case KindMnemonicAlgorand:
		return EncodeMnemonicAlgorand(src)
	default:
		return "", errUnsupportedEncoding
	}
}
