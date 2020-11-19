package util

import "encoding/base64"


type Base64Util struct {

}
func (base64Util Base64Util)Encoding(source []byte)  string {
	return base64.StdEncoding.EncodeToString(source)
}

func (base64Util Base64Util)Decoding(source string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(source)
}
