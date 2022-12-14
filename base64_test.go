package encoding

import "testing"

func TestEncodeBase64URL(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"alice-ed25519-key",
			args{
				[]byte{0x72, 0x3c, 0xaa, 0x23, 0xa5, 0xb5, 0x11, 0xaf, 0x5a, 0xd7, 0xb7, 0xef, 0x60, 0x76, 0xe4, 0x14, 0xab, 0x7e, 0x75, 0xa9, 0xdc, 0x91, 0xe, 0xa6, 0xe, 0x41, 0x7a, 0x2b, 0x77, 0xa, 0x56, 0x71},
			},
			"cjyqI6W1Ea9a17fvYHbkFKt-danckQ6mDkF6K3cKVnE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBase64URL(tt.args.src); got != tt.want {
				t.Errorf("EncodeBase64URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
