package encoding

func DecodeUTF8(src string) []byte {
	return []byte(src)
}

func EncodeUTF8(src []byte) string {
	return string(src)
}
