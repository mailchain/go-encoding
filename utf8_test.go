package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeUTF8(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"simple",
			args{
				[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 127, 125, 0, 1},
			},
			"\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\u007f}\x00\x01",
		},
		{
			"problem-sequence",
			args{
				[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 140, 125, 0, 1},
			},
			"\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\x8c}\x00\x01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, EncodeUTF8(tt.args.src))
		})
	}
}

func TestDecodeUTF8(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"simple",
			args{
				"\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\u007f}\x00\x01",
			},
			[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 127, 125, 0, 1},
		},
		{
			"problem-sequence",
			args{
				"\x00\x01\x02\x03\x04\x05\x06\a\b\t\n\x8c}\x00\x01",
			},
			[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 140, 125, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DecodeUTF8(tt.args.src))
		})
	}
}
