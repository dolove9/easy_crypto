package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadPrivateKey(t *testing.T) {
	assert := assert.New(t)
	keyUtil := KeyUtil{}

	keyBase64Encode, _ := FileUtil{}.ReadBytesFromFileName("C:/workspace/2021/dev/golang/resources/cert/dec_private.key")
	decryptedPriKey, _ := Base64Util{}.Decoding(string(keyBase64Encode))
	fmt.Println(FileUtil{}.DumpByteArrayToString(decryptedPriKey))

	privateKey, err := keyUtil.LoadPrivateKey(decryptedPriKey)

	if err != nil {
		fmt.Println(err)
	}

	err = privateKey.Validate()
	assert.NoError(err)

}
