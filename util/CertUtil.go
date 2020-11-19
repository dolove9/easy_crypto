package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"fmt"
	"math/big"
)

type CertUtil struct {
	certificate *x509.Certificate
}

func (certUtil *CertUtil)CertUtilFromFileName(path string)  (error) {
	fileUtil := FileUtil{}
	certByte, err := fileUtil.readBytesFromFileName(path)
	if err != nil{
		return err
	}


	certificate, err := x509.ParseCertificate(certByte)
	certUtil.certificate = certificate
	return err
}

func (certUtil CertUtil) CertBytes() [] byte  {
	return certUtil.certificate.Raw
}

func (certUtil CertUtil) SerialNumber() string {
	number := certUtil.certificate.SerialNumber
	newInt := big.NewInt(number.Int64())
	return newInt.String()
}

func (certUtil CertUtil) SerialNumberBigInt() big.Int {
	number := certUtil.certificate.SerialNumber
	return *number
}

func (certUtil CertUtil) SubjectDN() string {
	subject := certUtil.certificate.Subject
	return subject.String()
}

func (certUtil CertUtil)CrlDP() string {
	return certUtil.certificate.CRLDistributionPoints[0]
}

type VidResult struct {
	realName string
	digestAlg string
	vid [] byte
}
func (certUtil CertUtil)SubjectAltName()  (VidResult, error){
	extensions := certUtil.certificate.Extensions
	if extensions == nil{
		return VidResult{}, fmt.Errorf("extensions is not found")
	}

	identifier := asn1.ObjectIdentifier{2, 5, 29, 17} //SubjectAltName id
	//fmt.Println(len(extensions))
	for i := range extensions{
		if extensions[i].Id.Equal(identifier) { //SubjectAltName
			value := extensions[i].Value
			var temp asn1.RawValue
			asn1.Unmarshal(value, &temp)
			if temp.Tag == 16 { //SEQUENCE OF
				var temp2 asn1.RawValue
				asn1.Unmarshal(temp.Bytes, &temp2)

				var temp3 asn1.ObjectIdentifier //OID 1.2.410.200085
				rest, _ := asn1.Unmarshal(temp2.Bytes, &temp3)

				var temp4 asn1.RawValue	//CONTEXT_SPECIFIC
				asn1.Unmarshal(rest, &temp4)

				var temp5 asn1.RawValue
				asn1.Unmarshal(temp4.Bytes, &temp5) //SEQUENCE

				//fmt.Println(FileUtil{}.dumpByteArrayToString(temp5.Bytes))
				var realName asn1.RawValue
				r, _ := asn1.Unmarshal(temp5.Bytes, &realName) //UTF8String realName
				//fmt.Println(FileUtil{}.dumpByteArrayToString(realName.Bytes))
				//fmt.Println(string(realName.Bytes))
				//fmt.Println(FileUtil{}.dumpByteArrayToString(r))


				//vidSeq := asn1.ObjectIdentifier{ 1,2,410,200004,10,1,1,1}
				var vidSeq asn1.RawValue
				asn1.Unmarshal(r, &vidSeq)
				//fmt.Println("--" + strconv.Itoa(vidSeq.Tag))
				//fmt.Println("--" + FileUtil{}.dumpByteArrayToString(vidSeq.Bytes))

				var vidSeq2 asn1.RawValue
				asn1.Unmarshal(vidSeq.Bytes, &vidSeq2)
				var vidOid asn1.ObjectIdentifier
				vidStruct, _ := asn1.Unmarshal(vidSeq2.Bytes, &vidOid)
				//fmt.Println("--" + FileUtil{}.dumpByteArrayToString(vidStruct))

				//seq in vid struct
				var inVidSeq asn1.RawValue
				asn1.Unmarshal(vidStruct, &inVidSeq)
				//fmt.Println("--" + FileUtil{}.dumpByteArrayToString(inVidSeq.Bytes))

				var inVidSeq2 asn1.RawValue
				bytes, _ := asn1.Unmarshal(inVidSeq.Bytes, &inVidSeq2)
				//fmt.Println("--" + FileUtil{}.dumpByteArrayToString(inVidSeq2.Bytes))
				//fmt.Println("--" + FileUtil{}.dumpByteArrayToString(bytes))

				//HASH  ObjectIdentifier asn1.ObjectIdentifier{ 2,16,840,1,101,3,4,2,1} //SHA 256
				var vidBodyOctetString asn1.RawValue
				asn1.Unmarshal(bytes, &vidBodyOctetString)

				//OCTET_STRING
				var vidBody asn1.RawValue
				asn1.Unmarshal(vidBodyOctetString.Bytes, &vidBody)
				//fmt.Println("--" + FileUtil{}.dumpByteArrayToString(vidBody.Bytes))

				vidResult := new(VidResult)
				vidResult.realName = string(realName.Bytes)
				vidResult.digestAlg = vidOid.String()
				vidResult.vid = vidBody.Bytes
				return *vidResult, nil

			}
		}
	}
	return VidResult{}, fmt.Errorf("내부 오류가 발생하였습니다.")
}


func (certUtil CertUtil)PublicKey() *rsa.PublicKey {
	key, _ := x509.ParsePKIXPublicKey(certUtil.certificate.RawSubjectPublicKeyInfo)
	return key.(*rsa.PublicKey)
}