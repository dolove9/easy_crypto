package util

import (
	"crypto/sha1"
	"crypto/sha256"
	"errors"
)

type MDUtil struct {
	algName string
}

type digestResult struct {
	base64Result string
	digestResult [] byte
}

func (mdUtil *MDUtil) init(algName string) {
	mdUtil.algName = algName
}

func (mdUtil MDUtil) digest(source [] byte) (digestResult, error) {
	if mdUtil.algName == "" {
		return digestResult{}, errors.New("Alg Name is empty")
	}

	result := digestResult{}
	switch mdUtil.algName {
	case "SHA1":
		sum := sha1.Sum(source)
		result.digestResult = sum[:]
		result.base64Result = Base64Util{}.Encoding(sum[:])
	case "SHA256":
		sum256 := sha256.Sum256(source)
		result.digestResult = sum256[:]
		result.base64Result = Base64Util{}.Encoding(sum256[:])
	default:
		return digestResult{}, errors.New("지원하는 알고리즘이 없습니다.")

	}
	return result, nil
}
