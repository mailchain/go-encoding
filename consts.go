package encoding

// DataPrefix used to identify Mailchain messages.
func DataPrefix() []byte {
	return []byte{0x6d, 0x61, 0x69, 0x6c, 0x63, 0x68, 0x61, 0x69, 0x6e}
}

const (
	// KindHex encoding value.
	KindHex = "hex/plain"
	// KindHex0XPrefix encoding value.
	KindHex0XPrefix = "hex/0x-prefix"
	// KindBase32 encoding value.
	KindBase32 = "base32/plain"
	// KindBase58 encoding value.
	KindBase58 = "base58/plain"
	// KindBase64 encoding value.
	KindBase64 = "base64/plain"
	// KindBase64URL encoding value.
	KindBase64URL = "base64/url"
	// KindBase58SubstrateAddress encoding value.
	KindBase58SubstrateAddress = "base58/ss58-address"
	// KindUTF8 encoding value.
	KindUTF8 = "text/utf-8"
)
