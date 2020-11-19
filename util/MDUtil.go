package util

import "errors"

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

	return digestResult{}, nil
}
