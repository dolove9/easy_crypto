package util

import (
	"fmt"
	assert "github.com/stretchr/testify/assert"
	"testing"

)

func TestReadBytesFromFileName(t *testing.T)  {
	assert := assert.New(t)
	fileUtil := FileUtil{}
	name, _ := fileUtil.ReadBytesFromFileName("C:/workspace/2021/dev/golang/resources/cert/signCert.der")
	assert.Equal("", "")
	hexString := fileUtil.DumpByteArrayToString(name)
	fmt.Println(hexString)
}
