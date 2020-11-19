package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase64Util_Encoding(t *testing.T) {
	assert := assert.New(t)
	sourceString := "dolove9한글1234!@#"

	base64Util := Base64Util{}
	encoding := base64Util.Encoding([]byte(sourceString))
	assert.Equal("ZG9sb3ZlOe2VnOq4gDEyMzQhQCM=", encoding)
}

func TestBase64Util_Decoding(t *testing.T) {
	assert := assert.New(t)
	expectedString := "dolove9한글1234!@#"
	source := "ZG9sb3ZlOe2VnOq4gDEyMzQhQCM="

	base64Util := Base64Util{}
	decoding, err := base64Util.Decoding(source)
	assert.NoError(err)
	fmt.Println(string(decoding))
	assert.Equal(expectedString, string(decoding))
}