package libutils

import (
	"encoding/base64"
	"io"
)

type base64Util struct {
}

var Base64 = &base64Util{}

func (base *base64Util) EncodeStr(b []byte) string {
	// Base64 Standard Encoding
	sEnc := base64.StdEncoding.EncodeToString(b)
	return sEnc
}

func (base *base64Util) Encode(src []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	// Base64 Standard Encoding
	base64.StdEncoding.Encode(buf, src)
	return buf
}

func (base *base64Util) Decode(body io.ReadCloser) ([]byte, error) {
	decoder := base64.NewDecoder(base64.StdEncoding, body)
	defer body.Close()
	encrypted, err := io.ReadAll(decoder)
	return encrypted, err
}
