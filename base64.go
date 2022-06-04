// Copyright 2022 Mailchain Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encoding

import (
	"encoding/base64"
)

// DecodeBase64 returns the bytes represented by the base64 string src.
//
// DecodeBase64 expects that src contains only base64 characters.
// If the input is malformed, DecodeBase64 returns an error.
func DecodeBase64(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

// EncodeBase64 returns the string represented by the base64 byte src.
//
// EncodeBase64 expects that src contains only base64 byte.
// If the input is malformed, EncodeBase64 returns an error.
func EncodeBase64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// DecodeBase64URL returns the bytes represented by the base64 string src.
// This method decodes without padding.
func DecodeBase64URL(src string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(src)
}

// EncodeBase64URL returns the string represented by the base64 byte src.
// This method encodes without padding.
func EncodeBase64URL(src []byte) string {
	return base64.RawURLEncoding.EncodeToString(src)
}
