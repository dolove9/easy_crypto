package util

import (
	"crypto/rsa"
	"crypto/x509"
)

type KeyUtil struct {

}

func (keyUtil KeyUtil)LoadPrivateKey(decryptedPrivateKey []byte) (key *rsa.PrivateKey, err error) {
	//privateKey, err := x509.ParsePKCS1PrivateKey(decryptedPrivateKey)
	pkcs8PrivateKey, err := x509.ParsePKCS8PrivateKey(decryptedPrivateKey)

	privateKey := pkcs8PrivateKey.(*rsa.PrivateKey)

	return privateKey, err
}