package encoding

import (
	"testing"

	"github.com/mailchain/mailchain/testing/asserterror"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	type args struct {
		encoding string
		src      string
	}
	tests := []struct {
		name      string
		args      args
		want      []byte
		assertion assert.ErrorAssertionFunc
	}{
		{
			"0x-hex",
			args{
				"hex/0x-prefix",
				"0x5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761",
			},
			[]byte{0x56, 0x2, 0xea, 0x95, 0x54, 0xb, 0xee, 0x46, 0xd0, 0x3b, 0xa3, 0x35, 0xee, 0xd6, 0xf4, 0x9d, 0x11, 0x7e, 0xab, 0x95, 0xc8, 0xab, 0x8b, 0x71, 0xba, 0xe2, 0xcd, 0xd1, 0xe5, 0x64, 0xa7, 0x61},
			assert.NoError,
		},
		{
			"hex",
			args{
				"hex/plain",
				"5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761",
			},
			[]byte{0x56, 0x2, 0xea, 0x95, 0x54, 0xb, 0xee, 0x46, 0xd0, 0x3b, 0xa3, 0x35, 0xee, 0xd6, 0xf4, 0x9d, 0x11, 0x7e, 0xab, 0x95, 0xc8, 0xab, 0x8b, 0x71, 0xba, 0xe2, 0xcd, 0xd1, 0xe5, 0x64, 0xa7, 0x61},
			assert.NoError,
		},
		{
			"base58",
			args{
				"base58/plain",
				"5CLmNK8f16nagFeF2h3iNeeChaxPiAsJu7piNYJgdPpmaRzPD",
			},
			[]byte{0x9, 0x86, 0xc6, 0x71, 0x43, 0xad, 0x96, 0x6f, 0xa5, 0x79, 0xc9, 0x1b, 0x30, 0xc6, 0x7f, 0x95, 0xe7, 0x4b, 0xcc, 0xe3, 0xc5, 0xec, 0xb9, 0x5c, 0x96, 0xbf, 0xb5, 0x82, 0x87, 0x65, 0x64, 0xe4, 0x9c, 0x8, 0x3f, 0x1c},
			assert.NoError,
		},
		{
			"base32",
			args{
				"base32/plain",
				"BGDMM4KDVWLG7JLZZENTBRT7SXTUXTHDYXWLSXEWX62YFB3FMTSJYCB7DQ",
			},
			[]byte{0x9, 0x86, 0xc6, 0x71, 0x43, 0xad, 0x96, 0x6f, 0xa5, 0x79, 0xc9, 0x1b, 0x30, 0xc6, 0x7f, 0x95, 0xe7, 0x4b, 0xcc, 0xe3, 0xc5, 0xec, 0xb9, 0x5c, 0x96, 0xbf, 0xb5, 0x82, 0x87, 0x65, 0x64, 0xe4, 0x9c, 0x8, 0x3f, 0x1c},
			assert.NoError,
		},
		{
			"base64",
			args{
				"base64/plain",
				"wqaeV4smxgiFxiMXe3Du89yYwVf6w31Gzo5zMXc0fxs=",
			},
			[]byte{0xc2, 0xa6, 0x9e, 0x57, 0x8b, 0x26, 0xc6, 0x8, 0x85, 0xc6, 0x23, 0x17, 0x7b, 0x70, 0xee, 0xf3, 0xdc, 0x98, 0xc1, 0x57, 0xfa, 0xc3, 0x7d, 0x46, 0xce, 0x8e, 0x73, 0x31, 0x77, 0x34, 0x7f, 0x1b},
			assert.NoError,
		},
		{
			"base64-url",
			args{
				"base64/url",
				"wqaeV4smxgiFxiMXe3Du89yYwVf6w31Gzo5zMXc0fxs",
			},
			[]byte{0xc2, 0xa6, 0x9e, 0x57, 0x8b, 0x26, 0xc6, 0x8, 0x85, 0xc6, 0x23, 0x17, 0x7b, 0x70, 0xee, 0xf3, 0xdc, 0x98, 0xc1, 0x57, 0xfa, 0xc3, 0x7d, 0x46, 0xce, 0x8e, 0x73, 0x31, 0x77, 0x34, 0x7f, 0x1b},
			assert.NoError,
		},
		{
			"unknown",
			args{
				"unknown",
				"0x5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761",
			},
			nil,
			asserterror.EqualError(errUnsupportedEncoding),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.encoding, tt.args.src)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		encoding string
		src      []byte
	}
	tests := []struct {
		name      string
		args      args
		want      string
		assertion assert.ErrorAssertionFunc
	}{
		{
			"0x-hex",
			args{
				"hex/0x-prefix",
				[]byte{0x56, 0x2, 0xea, 0x95, 0x54, 0xb, 0xee, 0x46, 0xd0, 0x3b, 0xa3, 0x35, 0xee, 0xd6, 0xf4, 0x9d, 0x11, 0x7e, 0xab, 0x95, 0xc8, 0xab, 0x8b, 0x71, 0xba, 0xe2, 0xcd, 0xd1, 0xe5, 0x64, 0xa7, 0x61},
			},
			"0x5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761",
			assert.NoError,
		},
		{
			"hex",
			args{
				"hex/plain",
				[]byte{0x56, 0x2, 0xea, 0x95, 0x54, 0xb, 0xee, 0x46, 0xd0, 0x3b, 0xa3, 0x35, 0xee, 0xd6, 0xf4, 0x9d, 0x11, 0x7e, 0xab, 0x95, 0xc8, 0xab, 0x8b, 0x71, 0xba, 0xe2, 0xcd, 0xd1, 0xe5, 0x64, 0xa7, 0x61},
			},
			"5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761",
			assert.NoError,
		},
		{
			"base58",
			args{
				"base58/plain",
				[]byte{0x9, 0x86, 0xc6, 0x71, 0x43, 0xad, 0x96, 0x6f, 0xa5, 0x79, 0xc9, 0x1b, 0x30, 0xc6, 0x7f, 0x95, 0xe7, 0x4b, 0xcc, 0xe3, 0xc5, 0xec, 0xb9, 0x5c, 0x96, 0xbf, 0xb5, 0x82, 0x87, 0x65, 0x64, 0xe4, 0x9c, 0x8, 0x3f, 0x1c},
			},
			"5CLmNK8f16nagFeF2h3iNeeChaxPiAsJu7piNYJgdPpmaRzPD",
			assert.NoError,
		},
		{
			"base32",
			args{
				"base32/plain",
				[]byte{0x9, 0x86, 0xc6, 0x71, 0x43, 0xad, 0x96, 0x6f, 0xa5, 0x79, 0xc9, 0x1b, 0x30, 0xc6, 0x7f, 0x95, 0xe7, 0x4b, 0xcc, 0xe3, 0xc5, 0xec, 0xb9, 0x5c, 0x96, 0xbf, 0xb5, 0x82, 0x87, 0x65, 0x64, 0xe4, 0x9c, 0x8, 0x3f, 0x1c},
			},
			"BGDMM4KDVWLG7JLZZENTBRT7SXTUXTHDYXWLSXEWX62YFB3FMTSJYCB7DQ",
			assert.NoError,
		},
		{
			"base64",
			args{
				"base64/plain",
				[]byte{0xc2, 0xa6, 0x9e, 0x57, 0x8b, 0x26, 0xc6, 0x8, 0x85, 0xc6, 0x23, 0x17, 0x7b, 0x70, 0xee, 0xf3, 0xdc, 0x98, 0xc1, 0x57, 0xfa, 0xc3, 0x7d, 0x46, 0xce, 0x8e, 0x73, 0x31, 0x77, 0x34, 0x7f, 0x1b},
			},
			"wqaeV4smxgiFxiMXe3Du89yYwVf6w31Gzo5zMXc0fxs=",
			assert.NoError,
		},
		{
			"base64-url",
			args{
				"base64/url",
				[]byte{0xc2, 0xa6, 0x9e, 0x57, 0x8b, 0x26, 0xc6, 0x8, 0x85, 0xc6, 0x23, 0x17, 0x7b, 0x70, 0xee, 0xf3, 0xdc, 0x98, 0xc1, 0x57, 0xfa, 0xc3, 0x7d, 0x46, 0xce, 0x8e, 0x73, 0x31, 0x77, 0x34, 0x7f, 0x1b},
			},
			"wqaeV4smxgiFxiMXe3Du89yYwVf6w31Gzo5zMXc0fxs",
			assert.NoError,
		},
		{
			"unknown",
			args{
				"unknown",
				nil,
			},
			"",
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encode(tt.args.encoding, tt.args.src)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
